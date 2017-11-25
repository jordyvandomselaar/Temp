package main

import (
	"fmt"
	"github.com/jordyvandomselaar/Temp"
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

	view := Temp.View{
		Path: path.Join(currentScriptDir, "frontend/index.temp.html"),
	}

	fmt.Fprint(w, string(view.Parse()))
}
