package config

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
)

func DbSetup() (*sqlx.DB, error) {
	db, err := sql.Open("mysql", "arian:pes21@tcp(localhost:3307)/PcPologist")
	if err != nil {
		log.Errorf("db connection failed: %w", err)
		return nil, err
	}

	sqlxDB := sqlx.NewDb(db, "mysql")
	
	sqlxDB.SetMaxOpenConns(25)
	sqlxDB.SetMaxIdleConns(5)
	sqlxDB.SetConnMaxLifetime(5 * time.Minute)
	sqlxDB.SetConnMaxIdleTime(1 * time.Minute)
	
	err = sqlxDB.Ping()
	if err != nil {
		log.Debugf("db ping failed: %v", err)
	}

	return sqlxDB, nil
}