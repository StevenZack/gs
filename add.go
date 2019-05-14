package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func doAddRepo(confRepos []string) {
	reposToAdd := []string{}
	rs, e := parseAddLine(*flagAdd)
	if e != nil {
		fmt.Println("parse line error :", e)
		return
	}
	reposToAdd = append(reposToAdd, rs...)
	for _, arg := range flag.Args() {
		rs2, e := parseAddLine(arg)
		if e != nil {
			fmt.Println("parse args error :", e)
			continue
		}
		reposToAdd = append(reposToAdd, rs2...)
	}
	newConf := DistinctStringList(append(confRepos, reposToAdd...))
	for _, repo := range newConf {
		fmt.Println(color.YellowString("stored " + repo))
	}
	fmt.Println("all", len(newConf), "repos stored")
	e = writeConf(newConf)
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
