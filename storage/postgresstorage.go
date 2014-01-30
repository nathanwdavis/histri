package storage

import (
	"database/sql"
	"encoding/json"
	//"fmt"
	_ "github.com/lib/pq"
	"github.com/nathanwdavis/histri"
	"strconv"
	"time"
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
	//fmt.Println(newId)
	ev.Id = strconv.Itoa(newId)
	//fmt.Println(ev.Id)
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

func (self *PostgresStorage) ById(id string) (*histri.Event, error) {
	intId, err := idStrToInt(id)
	if err != nil {
		return nil, err
	}
	instance := new(histri.Event)
	result := self.conn.QueryRow(`select * from histri.events
								  where id = $1`, intId)
	if err = rowToEvent(result, instance); err != nil {
		return nil, err
	}
	return instance, nil
}

func (self *PostgresStorage) ByTimeRange(start, end time.Time) ([]histri.Event, error) {
	cursor, err := self.conn.Query(`select * from histri.events
									 where timeutc between $1 and $2`, start, end)
	if err != nil {
		return nil, err
	}
	events := make([]histri.Event, 0, 100)
	var instance histri.Event
	for cursor.Next() {
		if err = rowToEvent(cursor, &instance); err != nil {
			defer cursor.Close()
			return nil, err
		}
		events = append(events, instance)
	}
	return events, nil
}

func (self *PostgresStorage) connection() *sql.DB {
	return self.conn
}

func NewPostgresStorage() (*PostgresStorage, error) {
	db, err := sql.Open("postgres",
		"postgres://histri:postgres@127.0.0.1/event?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{
		db,
	}, nil
}

func idStrToInt(id string) (int64, error) {
	intId, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		return 0, err
	}
	return intId, nil
}

func rowToEvent(row rowScanner, event *histri.Event) error {
	var (
		intId     int64
		dataBytes []byte
	)
	err := row.Scan(&intId,
		&event.TimeUtc,
		&event.EventType,
		&event.ExtRef,
		&dataBytes)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(dataBytes, &event.Data); err != nil {
		return err
	}
	event.Id = strconv.Itoa(int(intId))
	return nil
}

type rowScanner interface {
	Scan(dest ...interface{}) error
}
