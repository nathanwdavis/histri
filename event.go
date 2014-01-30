package histri

import (
	//"fmt"
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

	t, err := ParseTimeStr(timeUtcStr)
	if err != nil {
		return nil, err
	}
	//fmt.Println(t)
	event := NewEvent(eventType, extRef, data, &t)
	//fmt.Println(event.TimeUtc)
	return event, nil
}

func ParseTimeStr(timeStr string) (time.Time, error) {
	t, retErr := time.Parse(time.RFC3339Nano, timeStr)
	//fmt.Println(t)
	if retErr != nil {
		if tym, err := time.Parse(time.RFC1123, timeStr); err == nil {
			t = tym
			//fmt.Println(t)
		} else if tym, err := time.Parse(time.UnixDate, timeStr); err == nil {
			t = tym
		} else {
			return t, retErr
		}
	}
	//fmt.Println(t)
	return t.UTC(), nil
}
