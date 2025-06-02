package logger

import (
	"os"
	"path/filepath"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// Log is the global logger instance
	Log *logrus.Logger
)

// Config holds the logger configuration
type Config struct {
	Environment string
	SentryDSN   string
}

// Initialize sets up the logger with the given configuration
func Initialize(config Config) error {
	Log = logrus.New()

	// Set log level
	Log.SetLevel(logrus.InfoLevel)

	// Create logs directory if it doesn't exist
	if err := os.MkdirAll("logs", 0755); err != nil {
		return err
	}

	// Configure file output with rotation
	fileRotate := &lumberjack.Logger{
		Filename:   filepath.Join("logs", "app.log"),
		MaxSize:    10, // MB
		MaxBackups: 30, // Keep 30 days of logs
		MaxAge:     30, // days
		Compress:   true,
	}

	// Set formatter based on environment
	if config.Environment == "production" {
		Log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	} else {
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	}

	// Set output to both file and stdout
	Log.SetOutput(fileRotate)

	// Initialize Sentry if DSN is provided and in production
	if config.SentryDSN != "" && config.Environment == "production" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:              config.SentryDSN,
			Environment:      config.Environment,
			TracesSampleRate: 1.0,
		})
		if err != nil {
			Log.WithError(err).Error("Failed to initialize Sentry")
			return err
		}
		Log.Info("Sentry initialized successfully")
	} else if config.Environment == "production" {
		Log.Warn("Sentry DSN not provided in production environment")
	}

	return nil
}

// CaptureError logs an error and sends it to Sentry if configured
func CaptureError(err error, fields map[string]interface{}) {
	if err == nil {
		return
	}

	// Add error to fields
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["error"] = err.Error()

	// Log to file
	Log.WithFields(fields).Error("Error occurred")

	// Send to Sentry if in production
	if sentry.CurrentHub().Client() != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			for k, v := range fields {
				scope.SetExtra(k, v)
			}
			sentry.CaptureException(err)
		})
	}
}

// Info logs an info message with optional fields
func Info(message string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	Log.WithFields(fields).Info(message)
}

// Error logs an error message with optional fields
func Error(message string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	Log.WithFields(fields).Error(message)
}

// Debug logs a debug message with optional fields
func Debug(message string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	Log.WithFields(fields).Debug(message)
}

// Warn logs a warning message with optional fields
func Warn(message string, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	Log.WithFields(fields).Warn(message)
}
