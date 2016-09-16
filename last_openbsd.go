// Copyright (c) 2016 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Package last implements parser of lastlog file on UNIX systems
package last

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/user"
	"strconv"
	"time"
)

const (
	// Log is system lastlog file
	Log       = "/var/log/lastlog"
	timeSize  = 8
	lineSize  = 8
	hostSize  = 256
	entrySize = timeSize + lineSize + hostSize
)

// Last login information
type Last struct {
	Time time.Time
	Line string
	Host string
}

func (l Last) String() string {
	return fmt.Sprintf("Last login: %v on %v from %v",
		l.Time.Format(time.UnixDate), l.Line, l.Host)
}

// Since retuns time passed sine last login
func (l Last) Since() time.Duration {
	return time.Since(l.Time)
}

// ByUID returns last system login of user by UID
func ByUID(uid int) (Last, error) {
	return FromFile(Log, uid)
}

// ByUser returns last system login of speciefed User
func ByUser(u *user.User) (Last, error) {
	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return Last{}, err
	}
	return FromFile(Log, uid)
}

// Current returns last login of current user
func Current() (Last, error) {
	cur, err := user.Current()
	if err != nil {
		return Last{}, err
	}
	return ByUser(cur)
}

// Username returns last login by username
func Username(name string) (Last, error) {
	u, err := user.Lookup(name)
	if err != nil {
		return Last{}, err
	}
	return ByUser(u)
}

// FromFile returns last login of user by UID from specified file
func FromFile(lastlog string, uid int) (Last, error) {
	f, err := os.Open(lastlog)
	if err != nil {
		return Last{}, err
	}
	defer f.Close()
	if _, err := f.Seek(int64(uid*entrySize), io.SeekStart); err != nil {
		return Last{}, err
	}
	buf := make([]byte, entrySize)
	if _, err := f.Read(buf); err != nil {
		return Last{}, err
	}
	t := binary.LittleEndian.Uint64(buf[:timeSize])
	if t == 0 {
		return Last{}, fmt.Errorf("Never logged in")
	}
	l := Last{
		Time: time.Unix(int64(t), 0),
		Line: trim(buf[timeSize : timeSize+lineSize]),
		Host: trim(buf[timeSize+lineSize:]),
	}
	return l, nil
}

func trim(b []byte) string {
	if i := bytes.IndexByte(b, 0); i >= 0 {
		b = b[:i]
	}
	return string(b)
}
