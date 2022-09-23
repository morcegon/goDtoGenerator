package dto

import "io/fs"

type filesToParse []fs.FileInfo

var Fp filesToParse

func (fp *filesToParse) AddFile(file fs.FileInfo) error {
	*fp = append(*fp, file)
	return nil
}
