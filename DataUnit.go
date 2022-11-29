package main

import "encoding/json"

type DataUnit struct {
	timeStamp int64
	FileName  string `json:"filename"`
	FileType  string `json:"type,omitempty"`
	Data      string `json:"data,omitempty"`
}

func deserializeDataUnit(data []byte) DataUnit {
	var o DataUnit
	ok := json.Unmarshal(data, &o)
	if ok != nil {
		panic(ok)
	}
	return o
}
