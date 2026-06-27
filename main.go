package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Found struct {
	Line  int16
	Cntnt string
}

type File struct {
	Name string
	Res  []Found
}

var (
	files = []File{}
)

func searchFile(current_dir string, word string) {
	entries, _ := os.ReadDir(current_dir)
	for _, entry := range entries {
		if entry.IsDir() {
			searchFile(current_dir+"/"+entry.Name(), word)
		} else {
			if strings.Contains(entry.Name(), ".") {
				data, _ := os.ReadFile(current_dir + "/" + entry.Name())
				if strings.Contains(string(data), word) {
					lines := strings.Split(string(data), "\n")
					foundlines := []Found{}
					for i, line := range lines {
						if strings.Contains(line, word) {
							foundlines = append(foundlines, Found{
								Line:  int16(i) + 1,
								Cntnt: line,
							})
						}
					}
					files = append(files, File{
						Name: entry.Name(),
						Res:  foundlines,
					})
				}
			}
		}
	}
}

func printing(files []File, color string, word string) {
	for _, file := range files {
		fmt.Println(file.Name)
		for _, val := range file.Res {
			Content := strings.ReplaceAll(val.Cntnt, word, color)
			fmt.Println(strconv.Itoa(int(val.Line)) + ":" + Content)
		}
		fmt.Print("\n")
	}
}

func main() {
	searchArr := os.Args[1:]
	current_dir := os.Getenv("PWD")
	//current_dir := "/home/coder/Projects/GoLearning/"
	searchFile(current_dir, searchArr[0])
	const reset = "\033[0m"
	const yellow = "\033[1;33m"
	Word := yellow + searchArr[0] + reset
	printing(files, Word, searchArr[0])
}
