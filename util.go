package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"

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

func flagParseWithArgs(f string) []string {
	return append(flag.Args(), f)
}

func stringListToMap(ss []string) map[string]bool {
	m := make(map[string]bool)
	for _, s := range ss {
		m[s] = true
	}
	return m
}

func mapToStringList(m map[string]bool) []string {
	ss := []string{}
	for k, _ := range m {
		ss = append(ss, k)
	}
	return ss
}

func filterGitRepoList(l []string) []string {
	out := []string{}
	for _, v := range l {
		path, e := checkIfIsGitRepo(v)
		if e != nil {
			continue
		}
		out = append(out, path)
	}
	return out
}

func checkIfIsGitRepo(dir string) (string, error) {
	_, e := gitstatusRepo(dir)
	if e != nil {
		return "", e
	}
	path, e := filepath.Abs(dir)
	if e != nil {
		return "", e
	}
	return path, nil
}
