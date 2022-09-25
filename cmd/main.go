package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/morcegon/goDtoGenerator/pkg/files"
)

func main() {
	var providedPath string
	flag.StringVar(&providedPath, "path", "", "")
	flag.Parse()

	if files.IsSingleFile(providedPath) {
		file, _ := os.Lstat(providedPath)
		files.FilesToParse.AddFile(file, providedPath)
	} else {
		files.ScanFolder(providedPath)
	}

	now := time.Now()
	for _, file := range files.FilesToParse {
		fmt.Println(file.Path)
	}

	elapsed := time.Since(now)
	fmt.Printf("\n\n%v files to be processed\n", len(files.FilesToParse))
	fmt.Printf("Spent %v", elapsed)
}
