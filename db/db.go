package db

import (
	"sync"
	"wg.dde1ta/entity"
)

type DB struct {
	Employees []entity.Employee `json:"employees"`
	Admins    []entity.Admin    `json:"admin"`
	dbFile    *DBFile           `json:"-"`
	isTransacting *sync.RWMutex
}


func NewDB(dbFilePath string) *DB {
	return &DB{
		Employees: make([]entity.Employee, 0),
		Admins: make([]entity.Admin, 0),
		dbFile: NewDBFile(dbFilePath),
	}
}

func NewDBBuffered(dbFilePath string, employeeCount int, adminCount int) *DB {
	return &DB{
		Employees: make([]entity.Employee, 0, employeeCount),
		Admins: make([]entity.Admin, 0, adminCount),
		dbFile: NewDBFile(dbFilePath),
	}
}



