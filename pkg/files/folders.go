package files

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

func ScanFolder(folderPath string) {
	files, _ := ioutil.ReadDir(folderPath)
	folderToScan := []fs.FileInfo{}

	for _, file := range files {
		if !file.IsDir() {
			Fp.AddFile(file)
		} else {
			folderToScan = append(folderToScan, file)
		}
	}

	for _, folder := range folderToScan {
		newFolderPath := filepath.Join(folderPath, folder.Name())
		ScanFolder(newFolderPath)
	}
}
