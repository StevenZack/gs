package main

import (
	"encoding/json"
	"log"
	"time"
)

var (
	repoMap = make(map[string]bool)
)

func init() {
	loadMap()
}

func loadMap() {
	e := json.Unmarshal([]byte(kvList.Get()), &repoMap)
	if e != nil {
		log.Fatal(e)
	}
}

func storeMap() {
	b, e := json.Marshal(repoMap)
	if e != nil {
		log.Fatal(e)
	}
	kvList.Post(string(b))
	time.Sleep(time.Millisecond * 20)
}
