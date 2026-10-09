package db

import (
	"encoding/json"
	"fmt"

	"wg.dde1ta/entity"
)

type DB struct {
	dbFile *DBFile `json:"-"`
}

type UserWrapper struct {
	User entity.User
}

type Seperator struct {
	Role string `json:"role"`
}

func NewDB(dbFilePath string) *DB {
	return &DB{
		// Employees: make([]entity.Employee, 0),
		// Admins: make([]entity.Admin, 0),
		dbFile: NewDBFile(dbFilePath),
	}
}

func (uw *UserWrapper) UnmarshalJSON(data []byte) error {
	// Extract only the "role" field
	var discriminator struct {
		Role string `json:"role"`
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return err
	}

	// Route based on the type
	switch discriminator.Role {
	case "employee":
		var emp entity.Employee
		if err := json.Unmarshal(data, &emp); err != nil {
			return err
		}
		uw.User = &emp
	case "admin":
		var admin entity.Admin
		if err := json.Unmarshal(data, &admin); err != nil {
			return err
		}
		uw.User = &admin
	default:
		return fmt.Errorf("unknown user type: %s", discriminator.Role)
	}

	return nil
}

func (db *DB) GetDB() ([]UserWrapper, error) {
	jsonRawSting, err := db.dbFile.fm.Read()

	if err != nil {
		return nil, err
	}

	var wrappers []UserWrapper

	err = json.Unmarshal([]byte(jsonRawSting), &wrappers)

	if err != nil {
		return nil, err
	}

	return wrappers, nil
}

func (db *DB) SaveToDB(data string) error {
	err := db.dbFile.fm.Write(data)

	return err
}
