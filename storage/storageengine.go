package storage

import (
	"errors"
	"github.com/nathanwdavis/histri"
	"strconv"
	"time"
)

// Interface that defines the ability to insert a new Event object
type Inserter interface {
	Insert(*histri.Event) error
}

// Interface that defines the ability to query for Events
type SimpleQuerier interface {
	ById(id string) (*histri.Event, error)
	ByTimeRange(start, end time.Time) ([]histri.Event, error)
	//ByExtRef(extRef string) ([]histri.Event, error)
}

type Counter interface {
	Count() (int64, error)
}

// Interface that defines a storage engine for Events
type Storage interface {
	Inserter
	SimpleQuerier
	Counter
}

// An implementation of Storage that should only be used for testing
type InMemStorage struct {
	events   []histri.Event
	isSorted bool
}

func (self *InMemStorage) Insert(event *histri.Event) error {
	event.Id = string(len(self.events))
	self.events = append(self.events, *event)
	return nil
}

func (self *InMemStorage) Count() (int64, error) {
	return int64(len(self.events)), nil
}

func (self *InMemStorage) ById(id string) (*histri.Event, error) {
	intId, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		return nil, err
	}
	return &self.events[intId-1], nil
}

func (self *InMemStorage) ByTimeRange(start, end time.Time) ([]histri.Event, error) {
	var results []histri.Event
	for _, ev := range self.events {
		if ev.TimeUtc.After(start) && ev.TimeUtc.Before(end) {
			results = append(results, ev)
		}
	}
	return results, nil
}

// Returns an implementation of Storage (currently only supports InMemStorage)
func NewStorage(typ string) (*Storage, error) {
	switch typ {
	case "", "postgres":
		postgresStore, err := NewPostgresStorage()
		if err != nil {
			return nil, err
		}
		inst := Storage(postgresStore)
		return &inst, nil
	case "inmem":
		inst := Storage(new(InMemStorage))
		return &inst, nil
	default:
		return nil, errors.New("Invalid storage type provided.")
	}
}
