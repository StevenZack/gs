package main

import (
	"encoding/json"

	"github.com/StevenZack/tools/fileToolkit"
)

func writeConf(c []string) error {
	f, e := fileToolkit.WriteFile(conf)
	if e != nil {
		return e
	}
	defer f.Close()
	b, e := json.Marshal(c)
	if e != nil {
		return e
	}
	f.Write(b)
	return nil
}

func readConf() ([]string, error) {
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
