package container

import (
	"fmt"
	"testing"
)

func TestGetServiceDir(t *testing.T) {
	beforeAll(t)

	path := GetServiceDir("customer")
	if path == "" {
		t.Fatal("path isn't get.")
	}

	fmt.Println(path)
}