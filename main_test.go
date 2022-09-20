package main

import "testing"

func TestIsFolder(t *testing.T) {
	folderPath := "dir1/dir2/"
	if !isFolder(folderPath) {
		t.Error("Expecting folder path")
	}

	folderPathWithoutSlash := "dir1/dir2"
	if !isFolder(folderPathWithoutSlash) {
		t.Error("Expecting folder path")
	}

	filenameWithExtension := "dir1/dir2/filename.ext"
	if isFolder(filenameWithExtension) {
		t.Error("Expecting file path")
	}
}
