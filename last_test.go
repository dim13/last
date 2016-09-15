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
	u, err := time.Parse(time.UnixDate, "Thu Sep 15 00:15:54 CEST 2016")
	if err != nil {
		t.Error(err)
	}
	if u != l.Time {
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
