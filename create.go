package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func adjustname(unit *DataUnit) {
	var newNum = 1
	FileBaseName := unit.FileName + "_"
	for {
		TempFileName := FileBaseName + strconv.FormatInt(int64(newNum), 10)
		if _, found := storage[TempFileName]; found {
			newNum++
		} else {
			log.Println("New name is... ", TempFileName)
			unit.FileName = TempFileName
			break
		}
	}
	return
}

func createFile(w http.ResponseWriter, r *http.Request) {
	// reading body
	bytes, ok := io.ReadAll(r.Body)
	if ok != nil {
		panic(ok)
	}
	// The whole part of adding a file to the storage
	data := deserializeDataUnit(bytes)
	data.timeStamp = time.Now().Unix()
	//checking if a file with that name exists
	if _, exists := storage[data.FileName]; exists {
		log.Println("Name already exists, a new name will be automatically assigned")
		adjustname(&data)
	}
	storage[data.FileName] = data
	log.Println("Stored file ", data.FileName)
}

/*
	just in case i decide to add those guards
	if r.Method != "POST" {
		w.WriteHeader(405)
		_, _ = fmt.Fprintf(w, "Method not supported")
		return
	}
*/
