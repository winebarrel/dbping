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
			db, err = config.OpenDB()

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
