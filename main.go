package main

import (
	"fmt"
	"net/http"
	"time"
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
func listFiles(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "The currently saved filenames and timestamps:\n")
	for k, v := range storage {
		fmt.Fprintf(
			w,
			"%v\t%v\t%v\n",
			time.Unix(v.timeStamp, 0),
			v.FileType,
			k,
		)
	}

}
