package histri

import (
	//"fmt"
	"testing"
	"time"
)

func TestNewEventWithTimeStr(t *testing.T) {
	timeString := "2013-12-05T23:50:06.080Z"
	event, err := NewEventWithTimeStr(
		timeString,
		"blah",
		"abc123",
		map[string]interface{}{
			"a": 1,
			"b": 2,
		},
	)

	if err != nil {
		t.Errorf("Time string parsing failed with error: %q", err.Error())
	}
	ns := event.TimeUtc.Nanosecond()
	expect_ns := 80000000
	if ns != expect_ns {
		t.Errorf("Milliseconds not parsed correctly. %d != %d", ns, expect_ns)
	}
}

//TODO: this is failing. I need to convert non-UTC times to UTC
func TestTimeZoneOffsetTimeString(t *testing.T) {
	timeString := "Fri, 06 Dec 2013 01:00:00 CST"
	event, err := NewEventWithTimeStr(
		timeString,
		"earthquake",
		"abc123",
		map[string]interface{}{
			"a": 1,
			"b": "b",
		},
	)

	if err != nil {
		t.Errorf("Time string parsing failed with error: %q", err.Error())
	}
	hour_utc := event.TimeUtc.Hour()
	expect_hour := 7
	if hour_utc != expect_hour {
		t.Errorf("Time not corrected to UTC. %d != %d", hour_utc, expect_hour)
	}
}

func TestInvalidRubyTimeString(t *testing.T) {
	timeString := "Mon Jan 02 15:04:05 -0700 2006"
	_, err := NewEventWithTimeStr(
		timeString,
		"moooo123",
		"98765432187653397836743298749233729857390",
		map[string]interface{}{
			"a":   1,
			"b":   2,
			"cat": false,
		},
	)

	if err == nil {
		t.Error("Error expected but was nil")
	}
	if pErr, ok := err.(*time.ParseError); !ok {
		t.Errorf("Invalid time string did return ParseError. Error: %q", pErr)
	}
}
