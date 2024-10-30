package admin

import (
	"net/http"

	"github.com/pkg/browser"
)

func Serve() {
	http.Handle("/", http.FileServer(http.Dir("admin/public/")))
	go func() {
		browser.OpenURL("http://localhost:5500")
	}()
	if err := http.ListenAndServe(":5500", nil); err != nil {
		panic(err)
	}
}