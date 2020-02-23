package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"

	"github.com/StevenZack/tools/strToolkit"
)

var (
	addRepo    = flag.String("a", "", "add single repo")
	removeRepo = flag.String("r", "", "remove repo")
	addDir     = flag.String("d", "", "add all repos in dir")
	list       = flag.Bool("l", false, "list all added repos")
	clear      = flag.Bool("c", false, "clear")
	verbose    = flag.Bool("v", false, "verbose")
)

func main() {
	flag.Parse()
	defer storeMap()
	if *addRepo != "" {
		path, e := doAddRepo(*addRepo)
		if e != nil {
			fmt.Println("add repo error :", e)
			return
		}
		if path != "" {
			fmt.Println(color.YellowString(path))
		}
		return
	}
	if *removeRepo != "" {
		path, e := doRemoveRepo(*removeRepo)
		if e != nil {
			fmt.Println("remove error :", e)
			return
		}
		if path != "" {
			fmt.Println(color.RedString(path))
		}
		return
	}
	if *addDir != "" {
		infos, e := ioutil.ReadDir(*addDir)
		if e != nil {
			fmt.Println("read dir error :", e)
			return
		}
		for _, info := range infos {
			path, e := doAddRepo(strToolkit.Getrpath(*addDir) + info.Name())
			if e != nil {
				continue
			}
			if path != "" {
				fmt.Println(color.YellowString(path))
			}
		}
		return
	}
	if *clear {
		repoMap = make(map[string]bool)
		return
	}
	if *list {
		for k := range repoMap {
			fmt.Println(k)
		}
		return
	}
}

func doRemoveRepo(dir string) (string, error) {
	path, e := checkGitDir(dir)
	if e != nil {
		return "", e
	}
	if _, ok := repoMap[path]; !ok {
		return "", nil
	}
	delete(repoMap, path)
	return path, nil
}

func doAddRepo(dir string) (string, error) {
	path, e := checkGitDir(dir)
	if e != nil {
		return "", e
	}
	if _, ok := repoMap[path]; ok {
		return "", nil
	}
	repoMap[path] = true
	return path, nil
}

func checkGitDir(dir string) (string, error) {
	path, e := filepath.Abs(dir)
	if e != nil {
		fmt.Println("abs error :", e)
		return "", e
	}
	info, e := os.Stat(path)
	if e != nil {
		fmt.Println("stat error :", e)
		return "", e
	}
	if !info.IsDir() {
		return "", errors.New(path + " is not dir")
	}
	_, e = checkIfIsGitRepo(dir)
	if e != nil {
		return "", e
	}
	return path, nil
}
