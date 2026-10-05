package files

import (
	"os"
	"path/filepath"
)

type FileManager struct {
	filePath string
	LockFile *LockFile
}

func NewFileManager(filepath string) *FileManager {
	return &FileManager{
		filePath: filepath,
		LockFile: NewLockFile(filepath + ".lock"),
	}
}

func (file FileManager) Read() (string, error) {
	err := file.LockFile.GetReadLock()

	if err != nil {
		return "", err
	}

	defer file.LockFile.UnlockReadLock()

	var fileContentByte []byte
	fileContentByte, err = os.ReadFile(file.filePath)
	if err != nil {
		return "", err
	}

	return string(fileContentByte), nil
}

func (file FileManager) Write(content string) error {

	dir := filepath.Dir(file.filePath)
	tempFile, err := os.CreateTemp(dir, "db-*.json.tmp")
	if err != nil {
		return err
	}
	tempPath := tempFile.Name()

	defer os.Remove(tempPath)

	_, err = tempFile.WriteString(content)
	
	tempFile.Close()
	if err != nil {
		return err
	}

	err = file.LockFile.GetWriteLock()
	if err != nil {
		return err
	}
	defer file.LockFile.UnlockWriteLock()

	err = os.Rename(tempPath, file.filePath)
	if err != nil {
		return err
	}

	return nil
}
