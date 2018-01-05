package timestamp

import (
	"errors"
	"time"
)

// TimeToTimestamp generates a millisecond timestamp from the time object.
func TimeToTimestamp(t time.Time) Timestamp {
	return Timestamp{
		TimeUnixMs: uint64(t.UnixNano() / 1000000),
	}
}

// TimestampToTime generates a time object from a millisecond timestamp.
func TimestampToTime(t uint64) time.Time {
	return time.Unix(0, int64(t)*1000000)
}

// ToTime converts the Timestamp to a time.Time
func (t *Timestamp) ToTime() time.Time {
	return TimestampToTime(t.TimeUnixMs)
}

// Now returns a timestamp for now.
func NowMs() Timestamp {
	return TimeToTimestamp(time.Now())
}

// Validate checks the timestamp.
func (t *Timestamp) Validate() error {
	if t == nil || t.TimeUnixMs == 0 {
		return errors.New("timestamp is empty")
	}
	return nil
}
