package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func readFile(w http.ResponseWriter, r *http.Request) {
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
	//serialise the found object
	serialised, ok := json.Marshal(storage[data.FileName])
	if ok != nil {
		panic(ok)
	}
	serialised = append(serialised, 10, 13)
	//write the data to the response for the client to decode
	_, err := w.Write(serialised)
	if err != nil {
		panic(err)
	}
	log.Println("sent file ", data.FileName)
}
