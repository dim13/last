// Copyright (c) 2016 Dimitri Sokolyuk. All rights reserved.
// Use of this source code is governed by ISC-style license
// that can be found in the LICENSE file.

// Package last implements parser of lastlog file on UNIX systems
package last

/*
#include <fcntl.h>
#include <utmp.h>
#include <unistd.h>

time_t
last(int uid)
{
	struct lastlog ll = {};
	off_t pos = (off_t)uid * sizeof(ll);
	int fd = open(_PATH_LASTLOG, O_RDONLY, 0);
	if (fd >= 0) {
		pread(fd, &ll, sizeof(ll), pos);
		close(fd);
	}
	return ll.ll_time;
}
*/
import "C"

import (
	"os/user"
	"strconv"
	"time"
)

// ByUID returns last system login of user by UID
func ByUID(uid int) (time.Time, error) {
	t := C.last(C.int(uid))
	return time.Unix(int64(t), 0), nil
}

// ByUser returns last system login of speciefed User
func ByUser(u *user.User) (time.Time, error) {
	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return time.Time{}, err
	}
	return ByUID(uid)
}

// Current returns last login of current user
func Current() (time.Time, error) {
	cur, err := user.Current()
	if err != nil {
		return time.Time{}, err
	}
	return ByUser(cur)
}

// Username returns last login by username
func Username(name string) (time.Time, error) {
	u, err := user.Lookup(name)
	if err != nil {
		return time.Time{}, err
	}
	return ByUser(u)
}
