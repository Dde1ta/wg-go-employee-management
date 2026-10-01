package files

import (
	"os"
	"path/filepath"
)

type FileManager struct {
	filePath string
	lockFile *LockFile
}

func NewFileManager(filepath string) *FileManager {
	return &FileManager{
		filePath: filepath,
		lockFile: NewLockFile(filepath + ".lock"),
	}
}

func (file FileManager) Read() (string, error) {
	err := file.lockFile.GetReadLock()

	if err != nil {
		return "", err
	}

	defer file.lockFile.UnlockReadLock()

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

	err = file.lockFile.GetWriteLock()
	if err != nil {
		return err
	}
	defer file.lockFile.UnlockWriteLock()

	err = os.Rename(tempPath, file.filePath)
	if err != nil {
		return err
	}

	return nil
}

