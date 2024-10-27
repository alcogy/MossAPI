package database

import (
	"testing"
)

func TestGetAllData(t *testing.T) {
	data := map[string]string{
		"customer": "12050",
		"hello": "12020",
		"representative": "12070",
		"project": "12060",
	}

	kvs := GetAllData()
	want := 4
	if len(kvs) != want {
		t.Fatalf("expected: %v, got: %v\n%v", want, len(kvs), kvs)
	}

	for _, v := range kvs {
		if data[v.Key] != v.Value {
			t.Fatalf("expected: %v, got: %v\n", data[v.Key], v.Value)
		}
	}
}