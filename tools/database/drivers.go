package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MariaDBConnection is a struct that implements the Connection interface for MariaDB/MySQL.
type MariaDBConnection struct{}

// PostgresConnection is a struct that implements the Connection interface for PostgreSQL.
type PostgresConnection struct{}

func PostgresDriver() *PostgresConnection {
	return &PostgresConnection{}
}

func MariaDBDriver() *MariaDBConnection {
	return &MariaDBConnection{}
}

// Connect establishes a connection to a MariaDB database using the provided configuration.
func (m *MariaDBConnection) Connect(cfg DatabaseCfg) (Conn, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return Conn{}, err
	}

	return Conn{Gorm: db}, nil
}

// / Connect establishes a connection to a PostgreSQL database using the provided configuration.
func (p *PostgresConnection) Connect(cfg DatabaseCfg) (Conn, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBDatabase, cfg.DBPort, cfg.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return Conn{}, err
	}

	return Conn{Gorm: db}, nil
}
