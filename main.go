package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/StevenZack/tools/fileToolkit"
)

var (
	sep         = string(os.PathSeparator)
	conf        = fileToolkit.GetHomeDir() + sep + ".gitstatus" + sep + "config.json"
	flagVerbose = flag.Bool("v", false, "verbose")
	flagAdd     = flag.String("a", "", "add repo to monitor")
	flagList    = flag.Bool("l", false, "list all currently monitored repos")
	flagClear   = flag.Bool("c", false, "clear all configure")
	flagRemove  = flag.String("r", "", "remove repo in monitor list")
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
		ls := flagParseWithArgs(*flagAdd)
		doAddRepo(repos, filterGitRepoList(ls)...)
		return
	}

	if *flagClear {
		doClear()
		return
	}

	if *flagList {
		doList(repos)
		return
	}

	if *flagRemove != "" {
		doRemove(repos, flagParseWithArgs(*flagRemove)...)
		return
	}

	for _, repo := range repos {
		e := parseStatusLine(repo)
		if e != nil {
			fmt.Println("parseRepoLine error :", e, "removing it")
			doRemove(repos, repo)
			continue
		}
	}
}
