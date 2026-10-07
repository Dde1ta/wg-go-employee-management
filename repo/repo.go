package repo

import (
	"wg.dde1ta/db"
	"wg.dde1ta/entity"
	"encoding/json"
)

type repo struct {
	db db.DB
}

func (r *repo) getUsers() ([]entity.User, error) {

	wapperArray, err := r.db.GetDB()

	if err != nil {
		return nil, err
	}

	var userSlice []entity.User = make([]entity.User, len(wapperArray))

	for idx, value := range wapperArray {
		userSlice[idx] = value.User
	}

	return userSlice, nil
}

func (r *repo) saveToDB(array []entity.User) error {
	toSave, err := json.Marshal(array)

	if err != nil {
		return err
	}

	err = r.db.SaveToDB(string(toSave))

	return err
}
