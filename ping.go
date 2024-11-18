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
			db, err = config.Open()

			if err != nil {
				continue
			}
		}

		for {
			now := time.Now()
			err = db.Ping()

			if err != nil {
				break
			}

			log.Printf("OK %s", time.Since(now))
			time.Sleep(time.Duration(config.Interval) * time.Second)
		}
	}
}
