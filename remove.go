package main

import (
	"encoding/json"
	"fmt"

	"github.com/StevenZack/tools/fileToolkit"
)

func doRemove(confRepos []string, rm ...string) {
	m := stringListToMap(confRepos)
	for _, r := range rm {
		delete(m, r)
	}
	ss := mapToStringList(m)
	f, e := fileToolkit.WriteFile(conf)
	if e != nil {
		fmt.Println("update conf error :", e)
		return
	}
	defer f.Close()
	b, e := json.Marshal(ss)
	if e != nil {
		fmt.Println("marshal error :", e)
		return
	}
	_, e = f.Write(b)
	if e != nil {
		fmt.Println("write error :", e)
		return
	}
	fmt.Println(rm, "removed from config list")
}
