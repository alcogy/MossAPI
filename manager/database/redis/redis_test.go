package redis

import (
	"os"
	"testing"
)

func beforeAll(t *testing.T) {
	cwd, _ := os.Getwd()
	t.Cleanup(func() {
		os.Chdir(cwd)
	})
	os.Chdir("../")
}

func TestFetchAllData(t *testing.T) {
	beforeAll(t)

	data := map[string]string{
		"customer": "12050",
		"hello": "12020",
		"representative": "12070",
		"project": "12060",
	}

	kvs := FetchAllData()
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
	beforeAll(t)

	if !CheckExistService("customer") {
		t.Fatalf("%v is exist.", "customer")
	}

	if CheckExistService("abc") {
		t.Fatalf("%v is not exist.", "abc")
	}
}

func TestCheckPortNumberFree(t *testing.T) {
	beforeAll(t)

	if CheckPortNumberFree("12050") {
		t.Fatalf("%v is not free.", "12050")
	}

	if !CheckPortNumberFree("13333") {
		t.Fatalf("%v is free.", "13333")
	}
}

func TestSetService(t *testing.T) {
	beforeAll(t)

	err := SetService("franc", "12100")
	if err != nil {
		t.Fatal(err)
	}

	kvs := FetchAllData()
	want := 5
	if len(kvs) != want {
		t.Fatalf("expected: %v, got: %v\n%v", want, len(kvs), kvs)
	}
}

func TestDeleteService(t *testing.T) {
	beforeAll(t)
	
	err := DeleteService("jone")
	if err != nil {
		t.Fatal(err)
	}

	kvs := FetchAllData()
	for _, v := range kvs {
		if v.Key == "jone" {
			t.Fatalf("jone is not deleted.")
		}
	}
}