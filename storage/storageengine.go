package storage

import (
	"github.com/nathanwdavis/histri"
	"time"
)

// Interface that defines the ability to insert a new Event object
type Inserter interface {
	Insert(*histri.Event) error
}

// Interface that defines the ability to query for Events
type SimpleQuerier interface {
	ById(id string) *histri.Event
	ByTimeRange(start, end time.Time) []histri.Event
	ByExtRef(extRef string) []histri.Event
}

type Counter interface {
	Count() (int64, error)
}

// Interface that defines a storage engine for Events
type Storage interface {
	Inserter
	//SimpleQuerier
	Counter
}

// An implementation of Storage that should only be used for testing
type InMemStorage struct {
	events   []histri.Event
	isSorted bool
}

func (self *InMemStorage) Insert(event *histri.Event) error {
	event.Id = string(len(self.events) + 1)
	self.events = append(self.events, *event)
	return nil
}

func (self *InMemStorage) Count() (int64, error) {
	return int64(len(self.events)), nil
}

// Returns an implementation of Storage (currently only supports InMemStorage)
func NewStorage(typ string) Storage {
	inst := Storage(new(InMemStorage))
	return inst
}
