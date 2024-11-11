package libs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFileTree(t *testing.T) {
	root := "./mytestdir"
	from, _ := filepath.Abs(root + "/fromdir")
	dist, _ := filepath.Abs(root + "/dist")

	os.Mkdir(root, 0750)
	os.Mkdir(from, 0750)
	os.Mkdir(from + "/subdir", 0750)
	os.WriteFile(from + "/test.txt", []byte("Hello world.\n"), 0755)
	os.WriteFile(from + "/subdir/sub.txt", []byte("Under world.\n"), 0755)
	
	os.Mkdir(dist, 0750)

	CopyFileTree(from, dist)

	info, err := os.Stat(dist + "/subdir")
	if err != nil || info == nil || !info.IsDir() {
		t.Fatal("None or error subdir")
	}

	info, err = os.Stat(dist + "/test.txt")
	if err != nil || info == nil || info.IsDir() {
		t.Fatal("None or error test.txt")
	}

	info, err = os.Stat(dist + "/subdir/sub.txt")
	if err != nil || info == nil || info.IsDir() {
		t.Fatal("None or error sub.txt")
	}

	path, _ := filepath.Abs(root)
	os.RemoveAll(path)
}