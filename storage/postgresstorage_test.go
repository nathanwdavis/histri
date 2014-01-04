package storage

import (
	//"fmt"
	"github.com/nathanwdavis/histri"
	"testing"
)

func TestNewPostgresStorage(t *testing.T) {
	_, err := NewPostgresStorage()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestPostresStorageInsert(t *testing.T) {
	db, err := NewPostgresStorage()
	if err != nil {
		t.Errorf("Failed to create new PostgresStorage. Error: %q", err.Error())
	}
	event1 := histri.NewEvent(
		"type",
		"abc123",
		map[string]interface{}{
			"a": 1,
			"b": 2,
		},
		nil,
	)
	event2 := histri.NewEvent(
		"type2",
		"abc123",
		map[string]interface{}{
			"a": 1,
			"c": 3,
		},
		nil,
	)
	event3 := histri.NewEvent(
		"type",
		"abc124",
		map[string]interface{}{
			"a": 1,
			"b": 2,
		},
		nil,
	)
	oldCount, err := db.Count()
	if err != nil {
		t.Errorf("Could not get Count. Error: %q", err.Error())
	}
	err = db.Insert(event1)
	if err != nil {
		t.Errorf("Could not Insert. Error: %q", err.Error())
	}
	db.Insert(event2)
	db.Insert(event3)

	if c, _ := db.Count(); c != oldCount+3 {
		t.Error("Count does not correctly reflect new Inserts.")
	}
}

func TestPostgresStorageInsertSetsNewId(t *testing.T) {
	db, _ := NewPostgresStorage()
	event := histri.NewEvent(
		"type",
		"abc123",
		map[string]interface{}{
			"a": 1,
			"b": 2,
		},
		nil,
	)
	oldId := event.Id
	db.Insert(event)
	if event.Id == oldId {
		t.Error("New Id was not set on Insert")
	}
}

func TestPostgresStorageInsertWithComplexJson(t *testing.T) {
	db, _ := NewPostgresStorage()
	event := histri.NewEvent(
		"type",
		"abc123",
		map[string]interface{}{
			"aaa": 9379534.84933,
			"bbb": false,
			"ccc": map[string]interface{}{
				"c_ddd": 123,
				"c_eee": []interface{}{"a", 123, true},
				"c_fff": []int64{389202384, 19890345, 8479893, -397234},
			},
			"ddd": struct {
				n int
				v string
			}{65, "abc123"},
		},
		nil,
	)
	err := db.Insert(event)
	if err != nil {
		t.Errorf("Could not Insert with complex json. Error: %q", err.Error())
	}
}

func TestPostgresStorageInsertWithTimeStr(t *testing.T) {
	db, _ := NewPostgresStorage()
	timeString := "Fri, 06 Dec 2013 01:00:00 CST"
	event, err := histri.NewEventWithTimeStr(
		timeString,
		"earthquake",
		"abc123",
		map[string]interface{}{
			"a": 1,
			"b": "b",
		},
	)
	//fmt.Println(event.TimeUtc)
	if err != nil {
		t.Error("Could not create Event with TimeStr")
	}
	err = db.Insert(event)
	if err != nil {
		t.Errorf("Could not Insert with time from string. Error: %q",
			err.Error())
	}
}
