package db

import (
	"context"
	"fmt"
	"groove-app/pkg/logger"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var Client *gorm.DB

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASS"),
		os.Getenv("PG_DB_NAME"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_SSL_MODE"),
		os.Getenv("PG_TZ"),
	)

	var err error
	if Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger.Default.LogMode(gormLogger.Silent)}); err != nil {
		logger.Fatal("connect postgres failed: %s", err.Error())
	}

	sqlDB, _ := Client.DB()

	// DB_MAX_OPEN_CONNS
	if maxOpen, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS")); err == nil {
		sqlDB.SetMaxOpenConns(maxOpen)
	} else {
		logger.Fatal("env DB_MAX_OPEN_CONNS is invalid: %s", err.Error())
	}

	// DB_MAX_CONN_LIFE_MINS
	if maxLife, err := strconv.Atoi(os.Getenv("DB_MAX_CONN_LIFE_MINS")); err == nil {
		sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(maxLife))
	} else {
		logger.Fatal("env DB_MAX_CONN_LIFE_MINS is invalid: %s", err.Error())
	}

	// DB_MAX_IDLE_CONNS
	if maxIdle, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS")); err == nil {
		sqlDB.SetMaxIdleConns(maxIdle)
	} else {
		logger.Fatal("env DB_MAX_IDLE_CONNS is invalid: %s", err.Error())
	}

	// DB_MAX_CONN_IDLE_MINS
	if maxIdleLife, err := strconv.Atoi(os.Getenv("DB_MAX_CONN_IDLE_MINS")); err == nil {
		sqlDB.SetConnMaxIdleTime(time.Minute * time.Duration(maxIdleLife))
	} else {
		logger.Fatal("env DB_MAX_CONN_IDLE_MINS is invalid: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err == nil {
		logger.Info("database client initialized")
	} else {
		logger.Fatal("failed to ping database: %s", err.Error())
	}
}
