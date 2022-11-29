package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func updateFile(w http.ResponseWriter, r *http.Request) {
	// reading body
	bytes, ok := io.ReadAll(r.Body)
	if ok != nil {
		panic(ok)
	}
	//deserializing the body
	data := deserializeDataUnit(bytes)
	// if a file does not exist then throw 404 not found
	if _, exists := storage[data.FileName]; !exists {
		w.WriteHeader(404)
		fmt.Fprintf(w, "No such file found\n")
		return
	}
	// updating the timestamp
	data.timeStamp = time.Now().Unix()
	//overwriting the previous value
	storage[data.FileName] = data
	log.Println("Updated file ", data.FileName)
}
