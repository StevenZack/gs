package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/StevenZack/tools/fileToolkit"
)

var (
	sep  = string(os.PathSeparator)
	conf = fileToolkit.GetHomeDir() + sep + ".gitstatus" + sep + "config.json"
)

func main() {
	b := gitstatusRepo(".")
	fmt.Println(b)
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

func gitstatusRepo(repo string) error {
	cmd := exec.Command("git", "status")
	reader, e := cmd.StdoutPipe()
	if e != nil {
		return e
	}
	e = cmd.Start()
	if e != nil {
		return e
	}
	b, e := ioutil.ReadAll(reader)
	if e != nil {
		return e
	}
	fmt.Println(string(b))
	return nil
}
