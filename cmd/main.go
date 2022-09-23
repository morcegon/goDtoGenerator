package main

import (
	"dtoGenerator/pkg/dto"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	var providedPath string
	flag.StringVar(&providedPath, "path", "", "")
	flag.Parse()

	if isSingleFile(providedPath) {
		file, _ := os.Lstat(providedPath)
		dto.Fp.AddFile(file)
	} else {
		scanFolder(providedPath)
	}

	fmt.Printf("%v files to be processed \n", len(dto.Fp))
	for _, file := range dto.Fp {
		fmt.Println(file.Name())
	}
}

func scanFolder(folderPath string) {
	files, _ := ioutil.ReadDir(folderPath)
	folderToScan := []fs.FileInfo{}

	for _, file := range files {
		if !file.IsDir() {
			dto.Fp.AddFile(file)
		} else {
			folderToScan = append(folderToScan, file)
		}
	}

	for _, folder := range folderToScan {
		newFolderPath := filepath.Join(folderPath, folder.Name())
		scanFolder(newFolderPath)
	}
}

func isSingleFile(filePath string) bool {
	fileInfo, err := os.Lstat(filePath)

	if err != nil {
		panic(err.Error())
	}

	return !fileInfo.IsDir()
}
