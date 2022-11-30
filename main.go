package main

import (
	"net/http"
)

var storage = make(map[string]DataUnit)

func main() {
	http.HandleFunc("/list", listFiles)
	http.HandleFunc("/create", createFile)
	http.HandleFunc("/read", readFile)
	http.HandleFunc("/delete", deleteFile)
	http.HandleFunc("/update", updateFile)

	ok := http.ListenAndServe("localhost:8080", nil)
	if ok != nil {
		panic(ok)
	}
}
