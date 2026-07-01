package logics

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

func searchFile(current_dir, word, name string, ignoreCase, ignoreDot, readGit bool, root string) {
	entries, _ := os.ReadDir(current_dir)
	for _, entry := range entries {
		if entry.IsDir() {
			if entry.Name() == ".git" && readGit == false {
				continue
			}
			name += entry.Name() + "/"
			searchFile(current_dir+"/"+entry.Name(), word, name, ignoreCase, ignoreDot, readGit, root)
			name = ""
		} else {
			if strings.Contains(entry.Name(), ".") {
				if root != "" && root != entry.Name() {
					continue
				}

				if ignoreDot && entry.Name()[0] == '.' {
					continue
				}

				data, _ := os.ReadFile(current_dir + "/" + entry.Name())
				content := string(data)
				if ignoreCase {
					word = strings.ToLower(word)
					content = strings.ToLower(content)
				}

				if strings.Contains(content, word) {
					lines := strings.Split(content, "\n")
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
						Name: "\033[1;32m" + name + entry.Name() + "\033[0m",
						Res:  foundlines,
					})
				}
			}
		}
	}
}

func printing(files []File, word string, color string) {
	for _, file := range files {
		fmt.Println(file.Name)
		for _, val := range file.Res {
			Content := strings.ReplaceAll(strings.ToLower(val.Cntnt), strings.ToLower(word), strings.ToLower(color))
			fmt.Println(strconv.Itoa(int(val.Line)) + ":" + Content)
		}
		fmt.Print("\n")
	}
}

type SearchArgs struct {
	Root              string
	SearchStr         string
	IngoreCase        bool
	IgnoreDotfile     bool
	ReadFromGitIgnore bool
	ReadFromGit       bool
}

func Search(
	searchArgs SearchArgs,
) {
	current_dir := os.Getenv("PWD")
	name := ""
	//current_dir := "/home/coder/Projects/GoLearning/"
	searchFile(current_dir, searchArgs.SearchStr, name, searchArgs.IngoreCase, searchArgs.IgnoreDotfile, searchArgs.ReadFromGit, searchArgs.Root)
	Word := "\033[1;41m" + searchArgs.SearchStr + "\033[0m"
	printing(files, searchArgs.SearchStr, Word)
}
