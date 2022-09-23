package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/morcegon/goDtoGenerator/pkg/files"
)

func main() {
	var providedPath string
	flag.StringVar(&providedPath, "path", "", "")
	flag.Parse()

	if files.IsSingleFile(providedPath) {
		file, _ := os.Lstat(providedPath)
		files.Fp.AddFile(file)
	} else {
		files.ScanFolder(providedPath)
	}

	fmt.Printf("%v files to be processed \n", len(files.Fp))
	for _, file := range files.Fp {
		fmt.Println(file.Name())
	}
}
