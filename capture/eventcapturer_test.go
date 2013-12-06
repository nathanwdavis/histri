package capture

import (
	"testing"
)

func TestCreateEvent(t *testing.T) {
	args := &CreateEventArgs{
		TimeUtc:   "2013-12-18",
		EventType: "step1",
		ExtRef:    "12345678",
		Data: map[string]interface{}{
			"a": 1,
			"b": 2,
		},
	}
	resp := new(string)

	capturer := new(Capturer)
	err := capturer.CreateEvent(args, resp)
	if err != nil {
		t.Errorf("CreateEvent failed with %q", err.Error())
	}
}
