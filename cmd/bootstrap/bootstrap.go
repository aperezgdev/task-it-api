package bootstrap

import (
	"io"
	"net/http"
)

func Run() error {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)

	return http.ListenAndServe(":3000", nil)
}