package main

import (
	"fmt"
	"net/http"
	"time"
)

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
