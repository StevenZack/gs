package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/fatih/color"
)

var (
	sep         = string(os.PathSeparator)
	conf        = fileToolkit.GetHomeDir() + sep + ".gitstatus" + sep + "config.json"
	flagVerbose = flag.Bool("v", false, "verbose")
	flagAdd     = flag.String("a", "", "add repo to monitor")
	flagList    = flag.Bool("l", false, "list current monitored repos")
)

func main() {
	flag.Parse()
	run()
}
func run() {
	repos, e := readConf()
	if e != nil {
		fmt.Println("read conf error :", e)
		return
	}
	if *flagAdd != "" { //  add mode
		doAddRepo(repos)
		return
	}

	if *flagList {
		for _, repo := range repos {
			fmt.Println(color.YellowString(repo))
		}
		if len(repos) == 0 {
			fmt.Println(color.RedString("no repo monitored for now"))
		} else {
			fmt.Println("all", len(repos), "repos stored")
		}
		fmt.Println("config file located at:", conf)
		return
	}

	for _, repo := range repos {
		e := parseStatusLine(repo)
		if e != nil {
			fmt.Println("parseRepoLine error :", e)
			return
		}
	}
}
