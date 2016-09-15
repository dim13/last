package last

import (
	"testing"
	"time"
)

func TestLast(t *testing.T) {
	l, err := FromFile("lastlog", 1000)
	if err != nil {
		t.Fatal(err)
	}
	u, err := time.ParseInLocation(time.UnixDate, "Thu Sep 14 22:15:54 UTC 2016", time.UTC)
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
	_, err := FromFile("lastlog", -1)
	if err == nil {
		t.Error("expected error")
	}
	_, err = FromFile("lastlog", 65535)
	if err == nil {
		t.Error("expected error")
	}
}
