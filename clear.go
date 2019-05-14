package main

import (
	"fmt"

	"github.com/StevenZack/tools/fileToolkit"
	"github.com/fatih/color"
)

func doClear() {
	f, e := fileToolkit.WriteFile(conf)
	if e != nil {
		fmt.Println("write config file error :", e)
		return
	}
	defer f.Close()
	_, e = f.WriteString("[]")
	if e != nil {
		fmt.Println("write error :", e)
		return
	}
	fmt.Println(color.GreenString("all cleared"))
}
