package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	id := 100
	if id != 100 {
		t.Fatal("Error")
	}
}
