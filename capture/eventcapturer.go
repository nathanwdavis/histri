package capture

import (
	"github.com/nathanwdavis/histri"
)

type Capturer struct {
	Event *histri.Event
}

type CreateEventArgs struct {
	TimeUtc   string                 `json:"timeUtc"`
	EventType string                 `json:"eventType"`
	ExtRef    string                 `json:"extRef"`
	Data      map[string]interface{} `json:"data"`
}

type CreateEventResponse struct {
	success bool
}

func (c *Capturer) CreateEvent(args *CreateEventArgs, resp *string) error {
	*resp = "8sfj03-83n9sj03-m9e3f9q2-t0mbd4e"
	return nil
}
