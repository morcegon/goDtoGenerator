package files

import (
	"io/fs"
	"os"
)

type filesToParse []fs.FileInfo

var Fp filesToParse

func (fp *filesToParse) AddFile(file fs.FileInfo) error {
	*fp = append(*fp, file)
	return nil
}

func IsSingleFile(filePath string) bool {
	fileInfo, err := os.Lstat(filePath)

	if err != nil {
		panic(err.Error())
	}

	return !fileInfo.IsDir()
}
