package db

import (
	"wg.dde1ta/files"
)

type DBFile struct {
	fm *files.FileManager
}

func NewDBFile(dbFilePath string) *DBFile{
	return &DBFile{
		fm: files.NewFileManager(dbFilePath),
	}
}
