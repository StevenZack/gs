package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/StevenZack/tools/strToolkit"
	"github.com/fatih/color"

	"github.com/StevenZack/tools/fileToolkit"
)

var (
	sep     = string(os.PathSeparator)
	conf    = fileToolkit.GetHomeDir() + sep + ".gitstatus" + sep + "config.json"
	verbose = flag.Bool("v", false, "verbose")
)

func main() {
	flag.Parse()
	parseRepoLine("/Users/stevenzacker/go/src/github.com/StevenZack/gitstatus")
}
func run() {
	_, e := findConf()
	if e != nil {
		fmt.Println("read conf error :", e)
		return
	}
}
func findConf() ([]string, error) {
	content, e := fileToolkit.ReadFileAll(conf)
	if e != nil {
		fo, e := fileToolkit.WriteFile(conf)
		if e != nil {
			return nil, e
		}
		defer fo.Close()
		fo.WriteString("[]")
		return nil, e
	}
	result := []string{}
	e = json.Unmarshal([]byte(content), &result)
	if e != nil {
		return nil, e
	}
	return result, nil
}

func parseRepoLine(repoLine string) error {
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
		if *verbose {
			fmt.Println(color.GreenString(repo))
		}
	} else {
		fmt.Println(color.RedString(repo))
	}
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
