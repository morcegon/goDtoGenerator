package files

import (
	"io/ioutil"
	"path"
)

func ScanFolder(folderPath string) {
	folderFiles, _ := ioutil.ReadDir(folderPath)
	arrayFolders := []string{}

	for _, fileInfo := range folderFiles {
		completePath := path.Join(folderPath, fileInfo.Name())

		if fileInfo.IsDir() {
			arrayFolders = append(arrayFolders, completePath)
			continue
		}

		FilesToParse.AddFile(fileInfo, completePath)
	}

	for _, folder := range arrayFolders {
		ScanFolder(folder)
	}
}
