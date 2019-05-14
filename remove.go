package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/fatih/color"
)

func doRemove(confRepos []string, rm ...string) {
	m := stringListToMap(confRepos)
	for _, r := range rm {
		path, e := filepath.Abs(r)
		if e != nil {
			fmt.Println("abs error :", e)
			continue
		}
		delete(m, path)
		defer fmt.Println(color.YellowString("removed " + path))
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
}
