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
		files.Fp.AddFile(file)
	} else {
		files.ScanFolder(providedPath)
	}

	now := time.Now()
	for _, file := range files.Fp {
		fmt.Println(file.Name())
	}

	elapsed := time.Since(now)
	fmt.Printf("\n\n%v files to be processed\n", len(files.Fp))
	fmt.Printf("Spent %v", elapsed)
}
