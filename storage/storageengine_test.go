package storage

import (
	"github.com/nathanwdavis/histri"
	"testing"
)

var (
	tStorage, _ = NewStorage("inmem")
	TStorage    = *tStorage
)

func TestStorageInsertSetsNewId(t *testing.T) {
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
	TStorage.Insert(event)
	if event.Id == oldId {
		t.Error("New Id was not set on Insert")
	}
}

func TestStorageInsertMultiple(t *testing.T) {
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
	oldCount, _ := TStorage.Count()
	TStorage.Insert(event1)
	TStorage.Insert(event2)
	TStorage.Insert(event3)

	if c, _ := TStorage.Count(); c != oldCount+3 {
		t.Error("Count does not correctly reflect new Inserts.")
	}
}
