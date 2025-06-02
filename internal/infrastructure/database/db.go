package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DBConnections holds all database connections
	DBConnections = make(map[string]*gorm.DB)
	RDB           *redis.Client
	mu            sync.RWMutex
)

// GetDB returns a database connection by name
func GetDB(name string) (*gorm.DB, error) {
	mu.RLock()
	defer mu.RUnlock()

	if db, exists := DBConnections[name]; exists {
		return db, nil
	}
	return nil, fmt.Errorf("database connection '%s' not found", name)
}

// InitDB initializes a database connection with the given name
func InitDB(name string) error {
	mu.Lock()
	defer mu.Unlock()

	// Check if connection already exists
	if _, exists := DBConnections[name]; exists {
		return fmt.Errorf("database connection '%s' already exists", name)
	}

	dsn := os.Getenv(fmt.Sprintf("DB_DSN_%s", name))
	if dsn == "" {
		return fmt.Errorf("environment variable DB_DSN_%s not set", name)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database %s: %v", name, err)
	}

	DBConnections[name] = db
	log.Printf("Database '%s' connected", name)
	return nil
}

// CloseDB closes a specific database connection
func CloseDB(name string) error {
	mu.Lock()
	defer mu.Unlock()

	if db, exists := DBConnections[name]; exists {
		sqlDB, err := db.DB()
		if err != nil {
			return fmt.Errorf("failed to get underlying *sql.DB for %s: %v", name, err)
		}
		if err := sqlDB.Close(); err != nil {
			return fmt.Errorf("failed to close database connection %s: %v", name, err)
		}
		delete(DBConnections, name)
		log.Printf("Database '%s' connection closed", name)
		return nil
	}
	return fmt.Errorf("database connection '%s' not found", name)
}

// CloseAllDBs closes all database connections
func CloseAllDBs() error {
	mu.Lock()
	defer mu.Unlock()

	for name, db := range DBConnections {
		sqlDB, err := db.DB()
		if err != nil {
			return fmt.Errorf("failed to get underlying *sql.DB for %s: %v", name, err)
		}
		if err := sqlDB.Close(); err != nil {
			return fmt.Errorf("failed to close database connection %s: %v", name, err)
		}
		delete(DBConnections, name)
		log.Printf("Database '%s' connection closed", name)
	}
	return nil
}

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	if err := RDB.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Redis connected successfully")
}

// InitAllDatabases initializes both CRM and IMS databases
func InitAllDatabases() error {
	// Initialize CRM database
	if err := InitDB("crm"); err != nil {
		return fmt.Errorf("failed to initialize CRM database: %v", err)
	}

	// Initialize IMS database
	// if err := InitDB("ims"); err != nil {
	// 	return fmt.Errorf("failed to initialize IMS database: %v", err)
	// }

	return nil
}
