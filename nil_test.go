package main

import (
	"testing"
	"time"
)

func init() {
	print("ssssssssss")
	// orm.Debug = true
}

type HoldsTime struct {
	time.Time
}

func TestBadAccess(t *testing.T) {
	var nilPointer *HoldsTime
	t.Log(nilPointer)
}
