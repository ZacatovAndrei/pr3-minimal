package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func deleteFile(w http.ResponseWriter, r *http.Request) {
	bytes, ok := io.ReadAll(r.Body)
	if ok != nil {
		panic(ok)
	}
	data := deserializeDataUnit(bytes)
	//if a file does not exist then throw 404 not found
	if _, exists := storage[data.FileName]; !exists {
		w.WriteHeader(404)
		fmt.Fprintf(w, "The file has not been found\n")
		return
	} else {
		delete(storage, data.FileName)
		log.Println("Deleted file ", data.FileName)
	}

}
