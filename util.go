package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/StevenZack/tools/strToolkit"
	"github.com/fatih/color"
)

func DistinctStringList(ss []string) []string {
	m := make(map[string]bool)
	for _, s := range ss {
		m[s] = true
	}
	result := []string{}
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}

func readDir(dir string) ([]string, error) {
	infos, e := ioutil.ReadDir(dir)
	if e != nil {
		return nil, e
	}
	strs := []string{}
	for _, info := range infos {
		if info.IsDir() {
			strs = append(strs, strToolkit.Getrpath(dir)+info.Name())
		}
	}
	return strs, nil
}

func flagParseWithArgs(f string) []string {
	return append(flag.Args(), f)
}

func stringListToMap(ss []string) map[string]bool {
	m := make(map[string]bool)
	for _, s := range ss {
		m[s] = true
	}
	return m
}

func mapToStringList(m map[string]bool) []string {
	ss := []string{}
	for k, _ := range m {
		ss = append(ss, k)
	}
	return ss
}

func filterGitRepoList(l []string) []string {
	out := []string{}
	for _, v := range l {
		path, e := checkIfIsGitRepo(v)
		if e != nil {
			continue
		}
		out = append(out, path)
	}
	return out
}

func checkIfIsGitRepo(dir string) (string, error) {
	_, e := gitstatusRepo(dir)
	if e != nil {
		return "", e
	}
	path, e := filepath.Abs(dir)
	if e != nil {
		return "", e
	}
	return path, nil
}

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
	path, e := filepath.Abs(repoLine)
	if e != nil {
		return e
	}
	handleBool(b, path)
	return nil
}

func handleBool(b bool, repo string) {
	if b {
		if *verbose {
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
	if strings.Contains(str, "git add") { //uncommited files remain
		return false, nil
	}
	return true, nil //nothing to commit
}
