package histri

import (
	"time"
)

type Event struct {
	TimeUtc   time.Time              `json:"timeUtc"`
	EventType string                 `json:"eventType"`
	ExtRef    string                 `json:"extRef"`
	Data      map[string]interface{} `json:"data"`
	Id        string                 `json:"id"`
}

// Creates a new Event object with TimeUtc defaulted to current UTC time.
func NewEvent(eventType, extRef string,
	data map[string]interface{},
	timeUtc *time.Time) *Event {

	if timeUtc == nil {
		utcNow := time.Now().UTC()
		timeUtc = &utcNow
	}

	return &Event{
		*timeUtc,
		eventType,
		extRef,
		data,
		"",
	}
}

// Creates a new Event object given a string for UTC time in RFC3339 format.
func NewEventWithTimeStr(timeUtcStr, eventType, extRef string,
	data map[string]interface{}) (*Event, error) {

	t, retErr := time.Parse(time.RFC3339Nano, timeUtcStr)
	if retErr != nil {
		if t, err := time.Parse(time.RFC1123, timeUtcStr); err == nil {
			t = t
		} else if t, err := time.Parse(time.UnixDate, timeUtcStr); err == nil {
			t = t
		} else {
			return nil, retErr
		}
	}
	t = t.UTC()
	event := NewEvent(eventType, extRef, data, &t)
	return event, nil
}
