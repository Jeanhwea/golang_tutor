package main

import (
	"testing"
	"unsafe"
)

func TestTutor05Pointer01(t *testing.T) {
	var nums1 int32 = -1
	// 任何类型的指针都可以转化成 unsafe.Pointer
	ptr := unsafe.Pointer(&nums1)

	// unsafe.Pointer 可以转化成任何类型的指针
	pUint := (*uint32)(ptr)
	pInt := (*int32)(ptr)

	t.Log(*pUint)
	t.Log(*pInt)
	t.Logf("%x\n", *pUint)
}
