package main

import (
	"fmt"
	"github.com/jordyvandomselaar/temp"
	"net/http"
	"os"
	"path"
)

func main() {
	http.HandleFunc("/", indexHandler)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	currentScriptDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	view := temp.View{
		Path: path.Join(currentScriptDir, "frontend/index.temp.html"),
	}

	html, err := view.Parse()

	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, html)
}
