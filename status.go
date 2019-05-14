package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func parseStatusLine(repoLine string) error {
	if strings.HasSuffix(repoLine, "*") {
		repos, e := readDir(repoLine[:len(repoLine)-1])
		if e != nil {
			return e
		}
		for _, repo := range repos {
			b, e := gitstatusRepo(repo)
			if e != nil {
				fmt.Println("status repo:"+repo+" error :", e)
				continue
			}
			handleBool(b, repo)
		}
		return nil
	}
	b, e := gitstatusRepo(repoLine)
	if e != nil {
		return e
	}
	handleBool(b, repoLine)
	return nil
}

func handleBool(b bool, repo string) {
	if b {
		if *flagVerbose {
			fmt.Println(color.GreenString(repo))
		}
	} else {
		fmt.Println(color.RedString(repo))
	}
}

func gitstatusRepo(repo string) (bool, error) {
	e := os.Chdir(repo)
	if e != nil {
		return false, e
	}

	cmd := exec.Command("git", "status")
	reader, e := cmd.StdoutPipe()
	if e != nil {
		return false, e
	}
	e = cmd.Start()
	if e != nil {
		return false, e
	}
	b, e := ioutil.ReadAll(reader)
	if e != nil {
		return false, e
	}
	str := string(b)
	if strings.Contains(str, ".git") {
		return false, errors.New(repo + " is not a git repo")
	}
	if strings.Contains(str, "git add") {
		return false, nil
	}
	return true, nil
}
