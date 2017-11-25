package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "<h1>Demo</h1>")
}
