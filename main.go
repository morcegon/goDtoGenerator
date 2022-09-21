package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	providedPath := "dtoFolder"
	filesToParse := []fs.FileInfo{}

	if isSingleFile(providedPath) {
		fileInfo, _ := os.Lstat(providedPath)
		filesToParse = append(filesToParse, fileInfo)
	} else {
		filesToParse = scanFolder(providedPath)
	}
}

func scanFolder(folderPath string) []fs.FileInfo {
	files, _ := ioutil.ReadDir(folderPath)
	returnFiles := []fs.FileInfo{}
	folderToScan := []fs.FileInfo{}

	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file.Name())
			returnFiles = append(returnFiles, file)
		} else {
			folderToScan = append(folderToScan, file)
		}
	}

	for _, folder := range folderToScan {
		newFolderPath := filepath.Join(folderPath, folder.Name())
		scanFolder(newFolderPath)
	}

	return returnFiles
}

func isSingleFile(filePath string) bool {
	fileInfo, err := os.Lstat(filePath)

	if err != nil {
		panic(err.Error())
	}

	return !fileInfo.IsDir()
}
