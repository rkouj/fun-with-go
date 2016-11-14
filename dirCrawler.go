package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"path/filepath"
	"time"
)

func main() {
	cwd := "/usr/local/google/home/rkouj/go/src/k8s.io/kubernetes"

	t := time.Now()
	filepath.Walk(cwd, count)
	fmt.Println(time.Since(t))
	key := "keyToSearch"
	t = time.Now()
	for _, file := range getFileAndDir(cwd) {
		doit(file, cwd, key)
	}
	fmt.Println(time.Since(t))
	fmt.Println(walk, mine)
	//doit(currentDir)
	// get files and dirs in current dir
	// iterate and if it's a dir, recurse
	// if it's a file, check for a particular value and update the file
	//getFileAndDir()
}

var walk = 0
var mine = 0

func count(path string, info os.FileInfo, err error) error {
	//fmt.Println(path, info.Name())
	walk++
	return nil
}

func doit(file os.FileInfo, cwd string, keyToSearch string) {
	//fmt.Println(file.Name(), cwd)
	mine++
	if file.IsDir() {
		for _, f := range getFileAndDir(cwd + "/" + file.Name()) {
			doit(f, cwd+"/"+file.Name(), keyToSearch)
		}

	} else {
		// Any opertaion on a file, for instance search for text in a file ?
		//	if content, _ := ioutil.ReadFile(cwd+"/"+file.Name()); strings.Contains(string(content), keyToSearch) {
		//fmt.Println("true", file.Name())
		//	}
	}
}

func getFileAndDir(dirName string) []os.FileInfo {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error kanda", err)
	}
	return files
}
