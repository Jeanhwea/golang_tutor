package code0512a

import (
	"testing"
)

func TestCode0512a01(t *testing.T) {
	t.Log(restoreIpAddress("1111"))
	t.Log(restoreIpAddress("000256"))
	t.Log(restoreIpAddress("25525522135"))
}
