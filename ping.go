package dbping

import (
	"database/sql"
	"log"
	"time"
)

func Ping(config *Config) {
	var db *sql.DB
	var err error

	for {
		if err != nil {
			log.Printf("[ERROR] %s", err)
			time.Sleep(time.Duration(config.Interval) * time.Second)
		}

		if db == nil {
			db, err = openDB(config)

			if err != nil {
				continue
			}
		}

		for {
			now := time.Now()
			v := "PING"

			if config.Query != "" {
				err = db.QueryRow(config.Query).Scan(&v)
			} else {
				err = db.Ping()
			}

			if err != nil {
				break
			}

			log.Printf("%s %s", v, time.Since(now))
			time.Sleep(time.Duration(config.Interval) * time.Second)
		}
	}
}

func openDB(config *Config) (*sql.DB, error) {
	db, err := sql.Open(string(config.Driver), config.DSN)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(0)
	db.SetConnMaxIdleTime(0)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)

	return db, nil
}
