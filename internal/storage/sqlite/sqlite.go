package sqlite

import (
	"database/sql"

	"github.com/sahasajib/students-api/internal/config"
	_"github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	Db *sql.DB

}

func New(cfg *config.Config)(*SQLiteStorage, error){
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	 id INTEGER PRIMARY KEY AUTOINCREMENT,
	 name TEXT NOT NULL,
	 email TEXT NOT NULL,
	 age INTEGER NOT NULL
	)`)
	if err != nil {
		return nil, err
	}
	return &SQLiteStorage{Db: db}, nil
}