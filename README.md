# Go CRM

A Gin boilerplate with multi-database support, event-driven architecture, and modern best practices.

## Features

- Gin Framework - Fast and efficient HTTP server
- Multi-Database Support - PostgreSQL with multiple database connections
- Event-Driven - Kafka integration for event handling
- Authentication - JWT-based authentication with Redis session store
- Structured Logging - Logrus with Sentry integration
- CLI Support - Cobra-based command-line interface
- Clean Architecture - Well-organized project structure
- API Documentation - Swagger/OpenAPI support (coming soon)

## Project Structure

The project follows the standard Go project layout with a focus on clean architecture principles:

```
.
├── cmd/                    # Main applications
│   └── api/               # API server entry point
├── internal/              # Private application code
│   ├── api/              # API layer
│   │   ├── handlers/     # HTTP handlers
│   │   ├── middleware/   # HTTP middleware
│   │   └── routes/       # Route definitions
│   ├── domain/           # Domain layer
│   │   ├── models/       # Domain models
│   │   └── repository/   # Repository interfaces
│   ├── service/          # Business logic layer
│   └── infrastructure/   # Infrastructure layer
│       ├── database/     # Database implementations
│       ├── kafka/        # Kafka implementations
│       └── config/       # Configuration
└── pkg/                  # Public library code
    └── utils/           # Public utility functions
```

### Structure Benefits

1. Clear Separation of Concerns
   - `cmd/`: Contains only the main application entry points
   - `internal/`: Contains all private application code
   - `pkg/`: Contains code that can be used by external applications

2. Domain-Driven Design
   - `domain/`: Contains core business logic and models
   - `service/`: Implements business logic using domain models
   - `repository/`: Defines interfaces for data access

3. Infrastructure Isolation
   - All external dependencies (database, kafka, etc.) are in `infrastructure/`
   - Easy to swap implementations without affecting business logic
   - Clear dependency direction: infrastructure → service → domain

4. API Layer Organization
   - `api/`: Contains all HTTP-related code
   - Clear separation between handlers, middleware, and routes
   - Easy to maintain and extend API endpoints

### Development Guidelines

When adding new features, follow these principles:

1. Domain Models
   - Place in `internal/domain/models`
   - Keep models pure and free of infrastructure concerns
   - Use interfaces to define contracts

2. Business Logic
   - Place in `internal/service`
   - Use domain models and repository interfaces
   - Keep services focused on business rules

3. API Handlers
   - Place in `internal/api/handlers`
   - Keep handlers thin, delegating to services
   - Use middleware for cross-cutting concerns

4. Infrastructure
   - Place in `internal/infrastructure`
   - Implement repository interfaces
   - Handle external service interactions

5. Public Utilities
   - Place in `pkg/` only if needed by external applications
   - Keep `internal/` for application-specific code

## Prerequisites

Before you begin, ensure you have the following installed:
- Go 1.24 or higher
- PostgreSQL 15 or higher
- Redis 7 or higher
- Kafka 3.x
- Zookeeper 3.x

## Installation

### 1. Database Setup

#### PostgreSQL
```bash
# Install PostgreSQL (if not already installed)
brew install postgresql@15

# Start PostgreSQL service
brew services start postgresql@15

# Create databases
createdb go_crm_crm
createdb go_crm_ims
```

### 2. Redis Setup
```bash
# Install Redis
brew install redis

# Start Redis service
brew services start redis
```

### 3. Kafka Setup
```bash
# Install Kafka
brew install kafka

# Start Zookeeper
brew services start zookeeper

# Start Kafka
brew services start kafka

# Create required Kafka topics
kafka-topics --create --topic user.created --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

### 4. Environment Setup

Create a `.env` file in the project root with the following variables:

```bash
# App Configuration
APP_ENV=development
PORT=8080
JWT_SECRET=your_jwt_secret_key

# Database Configuration
DB_DSN_crm=postgresql://localhost:5432/go_crm_crm?user=your_username&password=your_password
DB_DSN_ims=postgresql://localhost:5432/go_crm_ims?user=your_username&password=your_password

# Redis Configuration
REDIS_ADDR=localhost:6379

# Kafka Configuration
KAFKA_BROKERS=localhost:9092

# Sentry Configuration (Optional)
SENTRY_DSN=your_sentry_dsn
```

### 5. Project Setup

```bash
# Clone the repository
git clone https://github.com/yourusername/go-gin.git
cd go-gin

# Install dependencies
go mod download

# Run database migrations
go run main.go
```

## Running the Application

### Development Mode
```bash
# Run the server
go run main.go
```

### Production Mode
```bash
# Build the application
go build -o go-gin

# Run the application
./go-gin
```

## API Endpoints

### Public Routes
- `POST /register` - Register a new user
- `POST /login` - User login
- `GET /users` - Get all users

### Protected Routes
- `POST /api/tasks` - Create a new task
- `GET /api/tasks` - Get all tasks
- `PUT /api/tasks/:id` - Update a task
- `DELETE /api/tasks/:id` - Delete a task

## CLI Commands

```bash
# Send task reminders
go run main.go task:send-reminders
```


## Gin Framework Features

- Middleware Support
  - JWT Authentication
  - CORS
  - Request Logging
  - Error Handling

- Route Groups
  - Public Routes
  - Protected Routes
  - API Versioning

- Error Handling
  - Centralized Error Handling
  - Custom Error Responses
  - Validation Errors

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/xyz-feature`)
3. Commit your changes (`git commit -m 'Add some xyz feature'`)
4. Push to the branch (`git push origin feature/xyz-feature`)
5. Open a Pull Request
