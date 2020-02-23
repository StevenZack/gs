package main

import (
	"github.com/StevenZack/db"
	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/strToolkit"
)

var (
	appDir = strToolkit.Getrpath(fileToolkit.GetHomeDir()) + ".config/gs"
	mydb   = db.MustNewDB(appDir, "cypher")
)

var (
	kvList = mydb.String("list", "{}")
)
