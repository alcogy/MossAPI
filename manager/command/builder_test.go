package command

import (
	"fmt"
	"testing"
)

func TestExecuteBuildNonData(t *testing.T) {
	backend := readFile("C:\\Users\\info\\Dev\\MossAPI\\samples\\files\\build.json")
	if len(backend.Services) > 0 {
		t.Fatal("erorr")
	}

	if len(backend.Tables) > 0 {
		t.Fatal("erorr")
	}

	for _, v := range backend.Services {
		fmt.Println(v)
	}

	for _, v := range backend.Tables {
		fmt.Println(v)
	}
}

func TestExecuteBuild(t *testing.T) {
	backend := readFile("C:\\Users\\info\\Dev\\MossAPI\\samples\\files\\test2.json")
	if len(backend.Services) == 0 {
		t.Fatal("erorr")
	}

	if len(backend.Tables) == 0 {
		t.Fatal("erorr")
	}

	for _, v := range backend.Services {
		fmt.Println(v)
		if v.Service != "test1" {
			t.Fatal("erorr")
		}
	}

	for _, v := range backend.Tables {
		fmt.Println(v)
		if v.TableName != "mytable1" {
			t.Fatal("erorr")
		}
	}
}