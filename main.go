package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/StevenZack/tools/fileToolkit"
)

var (
	sep  = string(os.PathSeparator)
	conf = fileToolkit.GetHomeDir() + sep + ".gitstatus" + sep + "config.json"
)

func main() {
	fmt.Println(gitstatusRepo("/Users/stevenzacker/go/src/github.com/StevenZack/gitstatus"))
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

func parseRepoLine(repoLine string) {

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
	if strings.Contains(str, "git add") {
		return false, nil
	}
	return true, nil
}
