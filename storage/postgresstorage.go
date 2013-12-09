package storage

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/nathanwdavis/histri"
)

type PostgresStorage struct {
	conn *sql.DB
}

func (self *PostgresStorage) Insert(ev *histri.Event) error {
	newId := 0
	jsonData, err := json.Marshal(ev.Data)
	if err != nil {
		return err
	}
	row := self.conn.QueryRow(`select insert_event($1, $2, $3, $4)`,
		ev.TimeUtc,
		ev.EventType,
		ev.ExtRef,
		string(jsonData))
	if err := row.Scan(&newId); err != nil {
		return err
	}
	ev.Id = string(newId)
	return nil
}

func (self *PostgresStorage) Count() (int64, error) {
	var count int64
	result := self.conn.QueryRow(`select count(*) from histri.events`)
	if err := result.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func NewPostgresStorage() (Storage, error) {
	db, err := sql.Open("postgres",
		"postgres://histri:postgres@127.0.0.1/event?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return Storage(&PostgresStorage{
		db,
	}), nil
}
