package database

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"upnext-fng/tgpl-api/config"
	"upnext-fng/tgpl-api/pkg/logger"
)

// Connection represents the database connection
type Connection interface {
	DataSourceName() string
	Open() error
	Close() error
	Instance() (*gorm.DB, error)
	Ping() error
}

// connection is an implementation of the database Connection
type connection struct {
	dsn      string
	config   *Config
	instance *gorm.DB
}

// configureGormLogger configures the GORM logger based on the environment
func configureGormLogger(cfg config.Config) gormLogger.Interface {
	if cfg.IsDevelopment() {
		dbLogger := logger.NewLog("database", logger.WithDevelopment(cfg.IsDevelopment()))
		return logger.NewGormLogger(dbLogger, gormLogger.Config{
			SlowThreshold:             time.Second,     // Slow SQL threshold
			LogLevel:                  gormLogger.Info, // Log level
			IgnoreRecordNotFoundError: false,           // Not ignore not found error
			ParameterizedQueries:      false,           // Include params in the SQL log
			Colorful:                  true,            // Disable color
		})
	}

	// Production configuration - minimal logging
	dbLogger := logger.NewLog("database")
	return logger.NewGormLogger(dbLogger, gormLogger.Config{
		SlowThreshold:             time.Duration(200) * time.Millisecond, // 200ms threshold
		LogLevel:                  gormLogger.Warn,                       // Only warnings and errors
		IgnoreRecordNotFoundError: true,                                  // Ignore record not found errors
		ParameterizedQueries:      true,                                  // Show parameterized queries
		Colorful:                  false,                                 // No color in production
	})
}

// NewConnection creates and returns a database connection instance
func NewConnection(dsn string, config *Config) Connection {
	if config == nil {
		config = newDefaultConfig()
	}

	if config.Config == nil {
		config.Config = newGormConfig()
	}

	return &connection{dsn: dsn, config: config}
}

// DataSourceName returns the data source connection string
func (c *connection) DataSourceName() string {
	return c.dsn
}

// Open initializes a new database client
func (c *connection) Open() error {
	if c.config == nil || c.config.Config == nil {
		return ErrMissingConfig
	}

	var err error
	c.instance, err = gorm.Open(mysql.Open(c.dsn), c.config.Config)
	if nil != err {
		return err
	}

	instanceDb, err := c.instance.DB()
	if nil != err {
		return err
	}

	if c.config.MaxOpenConnections > 0 {
		instanceDb.SetMaxOpenConns(c.config.MaxOpenConnections)
	}

	if c.config.MaxIdleConnections > 0 {
		instanceDb.SetMaxIdleConns(c.config.MaxIdleConnections)
	}

	if c.config.ConnectionMaxTime > 0 {
		instanceDb.SetConnMaxLifetime(c.config.ConnectionMaxTime)
	}

	if c.config.ConnectionIdleTime > 0 {
		instanceDb.SetConnMaxIdleTime(c.config.ConnectionIdleTime)
	}

	return nil
}

// Close closes the current database client
func (c *connection) Close() error {
	if c.instance == nil {
		return ErrUninitializedDatabase
	}

	gormDb, err := c.instance.DB()
	if err != nil {
		return err
	}

	return gormDb.Close()
}

// Instance return the current instance of the Mongo database client
func (c *connection) Instance() (*gorm.DB, error) {
	if c.instance == nil {
		return nil, ErrUninitializedDatabase
	}

	return c.instance, nil
}

// Ping verifies if the current database client is active and healthy
func (c *connection) Ping() error {
	instance, err := c.Instance()
	if err != nil {
		return err
	}

	gormDb, err := instance.DB()
	if err != nil {
		return err
	}

	return gormDb.Ping()
}
