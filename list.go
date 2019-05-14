package main

import (
	"fmt"

	"github.com/fatih/color"
)

func doList(repos []string) {
	for _, repo := range repos {
		fmt.Println(color.YellowString(repo))
	}
	if len(repos) == 0 {
		fmt.Println(color.RedString("no repo monitored for now"))
	} else {
		fmt.Println("all", len(repos), "repos stored")
	}
	fmt.Println("config file located at:", conf)
}
