package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func doAddRepo(confRepos []string, add ...string) {
	newConf := DistinctStringList(append(confRepos, add...))
	for _, repo := range newConf {
		fmt.Println(color.YellowString("stored " + repo))
	}
	fmt.Println("all", len(newConf), "repos stored")
	e := writeConf(newConf)
	if e != nil {
		fmt.Println("write conf error :", e)
		return
	}
}

func parseAddLine(line string) ([]string, error) {
	if strings.HasSuffix(line, "*") {
		lines, e := readDir(line[:len(line)-1])
		if e != nil {
			return nil, e
		}
		return lines, nil
	}
	p, e := filepath.Abs(line)
	if e != nil {
		return nil, e
	}
	_, e = gitstatusRepo(p)
	if e != nil {
		return nil, e
	}
	return []string{p}, nil
}
