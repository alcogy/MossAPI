package libs

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func CopyFileTree(root string, dist string) {
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		fmt.Println(path)

		// root := "../services/" + flags.Service
		from := strings.Replace(path, root, "", -1)

		// Make directory
		if info.IsDir() && path != dist {
			dist := dist + from
			os.MkdirAll(dist, 0750)
		}

		if !info.IsDir() {
			r, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer r.Close()

			w, err := os.Create(dist + from)
			if err != nil {
				panic(err)
			}
			defer w.Close()

			io.Copy(w, r)
		}
		return nil
	})
}