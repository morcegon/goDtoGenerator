package main

import "path"

type File struct{}
type Folder struct{}

func main() {
	if isFolder("dtoFolder/testDto.json") {

	}
}

func isFolder(filePath string) bool {
	extension := path.Ext(filePath)
	return len(extension) == 0
}
