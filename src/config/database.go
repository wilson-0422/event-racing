package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", DBPath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	DB.SetMaxOpenConns(1)
	DB.SetMaxIdleConns(1)
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	migrate()
	log.Println("Database initialized successfully")
}

func migrate() {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			role TEXT DEFAULT 'user',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS athletes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			gender TEXT NOT NULL,
			age INTEGER,
			team TEXT,
			event TEXT,
			phone TEXT,
			id_number TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS groups (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			event TEXT NOT NULL,
			status TEXT DEFAULT 'pending',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS group_athletes (
			group_id INTEGER NOT NULL,
			athlete_id INTEGER NOT NULL,
			PRIMARY KEY (group_id, athlete_id),
			FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
			FOREIGN KEY (athlete_id) REFERENCES athletes(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS venues (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			location TEXT,
			capacity INTEGER,
			status TEXT DEFAULT 'available',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS schedules (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			group_id INTEGER NOT NULL,
			venue_id INTEGER NOT NULL,
			start_time DATETIME NOT NULL,
			end_time DATETIME,
			status TEXT DEFAULT 'scheduled',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
			FOREIGN KEY (venue_id) REFERENCES venues(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS scores (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			athlete_id INTEGER NOT NULL,
			schedule_id INTEGER NOT NULL,
			score TEXT,
			rank INTEGER,
			remark TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (athlete_id) REFERENCES athletes(id) ON DELETE CASCADE,
			FOREIGN KEY (schedule_id) REFERENCES schedules(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS awards (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			athlete_id INTEGER NOT NULL,
			event TEXT NOT NULL,
			medal_type TEXT NOT NULL,
			competition_name TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (athlete_id) REFERENCES athletes(id) ON DELETE CASCADE
		)`,
	}
	for _, stmt := range statements {
		if _, err := DB.Exec(stmt); err != nil {
			log.Fatal("Migration failed:", err)
		}
	}
}
