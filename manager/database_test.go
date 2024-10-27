// TODO Move file place to correct.
package main

import (
	"manager/database/redis"
	"testing"
)

func TestFetchAllData(t *testing.T) {
	data := map[string]string{
		"customer": "12050",
		"hello": "12020",
		"representative": "12070",
		"project": "12060",
	}

	kvs := redis.FetchAllData()
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

func TestCheckExistService(t *testing.T) {
	if !redis.CheckExistService("customer") {
		t.Fatalf("%v is exist.", "customer")
	}

	if redis.CheckExistService("abc") {
		t.Fatalf("%v is not exist.", "abc")
	}
}

func TestCheckPortNumberFree(t *testing.T) {
	if redis.CheckPortNumberFree("12050") {
		t.Fatalf("%v is not free.", "12050")
	}

	if !redis.CheckPortNumberFree("13333") {
		t.Fatalf("%v is free.", "13333")
	}
}

func TestSetService(t *testing.T) {
	err := redis.SetService("jone", "120900")
	if err != nil {
		t.Fatal(err)
	}

	kvs := redis.FetchAllData()
	want := 5
	if len(kvs) != want {
		t.Fatalf("expected: %v, got: %v\n%v", want, len(kvs), kvs)
	}
}

func TestDeleteService(t *testing.T) {
	err := redis.DeleteService("jone")
	if err != nil {
		t.Fatal(err)
	}

	kvs := redis.FetchAllData()
	for _, v := range kvs {
		if v.Key == "jone" {
			t.Fatalf("jone is not deleted.")
		}
	}
}