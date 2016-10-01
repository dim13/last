package last

import (
	"testing"
	"time"
)

func TestLast(t *testing.T) {
	l, err := FromFile(testLog, testUID)
	if err != nil {
		t.Fatal(err)
	}
	u, err := time.ParseInLocation(time.UnixDate, testTime, time.UTC)
	if err != nil {
		t.Error(err)
	}
	if u != l.Time.In(time.UTC) {
		t.Errorf("expected %v, got %v", u, l.Time)
	}
	t.Log(l)
	t.Log(l.Since())
}

func TestSeek(t *testing.T) {
	_, err := FromFile(testLog, -1)
	if err == nil {
		t.Error("expected error")
	}
	_, err = FromFile(testLog, 65535)
	if err == nil {
		t.Error("expected error")
	}
}
