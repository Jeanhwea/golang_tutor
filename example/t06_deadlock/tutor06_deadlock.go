package main

import (
	"bytes"
	"runtime"
	"strconv"
)

func GetGid() (gid int64) {
	var buf [64]byte
	s := buf[:runtime.Stack(buf[:], false)]
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	gid, _ = strconv.ParseInt(string(s), 10, 64)
	return
}
