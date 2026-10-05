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


func (dbFile *DBFile) StartWrite() error {
	return dbFile.fm.LockFile.GetWriteLock()
}

func (dbFile *DBFile) CompleteWrite() error {
	return dbFile.fm.LockFile.UnlockWriteLock()
}
