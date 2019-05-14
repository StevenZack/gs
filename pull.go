package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/fatih/color"
)

func doPull(repos []string) {
	wg := new(sync.WaitGroup)
	wg.Add(len(repos))
	for _, repo := range repos {
		go doSinglePull(repo, wg)
	}
	wg.Wait()
}
func doSinglePull(repo string, wg *sync.WaitGroup) {
	defer wg.Done()
	//pull
	b, e := singlePull(repo)
	if e != nil {
		fmt.Println("single pull error :", e)
		return
	}
	if b {
		fmt.Println(color.BlueString("pulled " + repo))
	} else {
		fmt.Println(color.GreenString(repo))
	}
}
func singlePull(repo string) (bool, error) {
	e := os.Chdir(repo)
	if e != nil {
		return false, e
	}
	cmd := exec.Command("git", "pull", "origin", "master")
	pip, e := cmd.StdoutPipe()
	if e != nil {
		return false, e
	}
	e = cmd.Start()
	if e != nil {
		return false, e
	}
	b, e := ioutil.ReadAll(pip)
	if e != nil {
		return false, e
	}
	str := string(b)
	if strings.Contains(str, "+") || strings.Contains(str, "-") {
		return true, nil
	}
	return false, nil
}
