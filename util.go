package main

import (
	"io/ioutil"

	"github.com/StevenZack/tools/strToolkit"
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
