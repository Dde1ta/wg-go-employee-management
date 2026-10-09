package files

import (
	"github.com/gofrs/flock"
	// "time"
)

type LockFile struct {
	lockFile string
	fileLock *flock.Flock
}

func NewLockFile(lockFilePath string) *LockFile {
	return &LockFile{
		lockFile: lockFilePath,
		fileLock: flock.New(lockFilePath),
	}
}

func (lock *LockFile) GetWriteLock() error {
	return lock.fileLock.Lock()
}

func (lock *LockFile) UnlockWriteLock() error {
	return lock.fileLock.Unlock()
}

func (lock *LockFile) GetReadLock() error {

	return lock.fileLock.RLock()
}

func (lock *LockFile) UnlockReadLock() error {
	return lock.fileLock.Unlock()
}
