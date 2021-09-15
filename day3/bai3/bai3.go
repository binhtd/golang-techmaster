package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	directory := getInputData()
	ignoreDirectories := []string{".git"}
	err := filepath.Walk(directory, printFile(directory, ignoreDirectories))

	if err != nil {
		log.Println(err)
	}
}

/*
get user input directory info
*/
func getInputData() string {
	if len(os.Args) > 1 {
		return strings.TrimSuffix(os.Args[1], "/")
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	return strings.TrimSuffix(dir, "/")
}

/*
print each entry in tree
*/
func printFile(directory string, ignoreDirs []string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if info.IsDir() {
			dir := filepath.Base(path)
			for _, d := range ignoreDirs {
				if d == dir {
					return filepath.SkipDir
				}
			}
		}

		fmt.Println(formatPrintingPath(directory, path))
		return nil
	}
}

/*
format each entry in tree
*/
func formatPrintingPath(directory string, path string) string {
	formatPath := removeTopDirectoryAndBeginSlashAndEndingSlash(directory, path)
	if isTopTreeDirectory(formatPath) {
		return filepath.Base(path)
	}

	nestedLevel := strings.Count(formatPath, "/")

	return strings.Repeat("│   ", nestedLevel-1) + fmt.Sprintf("├── %s", filepath.Base(formatPath))
}

/*
remove top directory and slash
*/
func removeTopDirectoryAndBeginSlashAndEndingSlash(directory string, path string) string {
	path = strings.TrimPrefix(path, directory)
	path = strings.TrimSuffix(path, "/")

	return path
}

/*
checking entry is top directory or not
*/
func isTopTreeDirectory(path string) bool {
	return path == ""
}
