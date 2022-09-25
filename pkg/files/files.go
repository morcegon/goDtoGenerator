package files

import (
	"io/fs"
	"os"
	"path/filepath"
)

type FileToParse struct {
	FileInfo fs.FileInfo
	Path     string
}

type arrayFilesToParse []FileToParse

var FilesToParse arrayFilesToParse

func (fp *arrayFilesToParse) AddFile(fileInfo fs.FileInfo, path string) error {
	expectedExtension := ".json"
	if filepath.Ext(path) == expectedExtension {
		file := FileToParse{
			fileInfo,
			path,
		}

		*fp = append(*fp, file)
	}

	return nil
}

func IsSingleFile(filePath string) bool {
	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		panic(err.Error())
	}

	return !fileInfo.IsDir()
}
