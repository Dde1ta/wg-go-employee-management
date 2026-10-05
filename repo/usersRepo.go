package repo


import (
	"wg.dde1ta/entity"
	"wg.dde1ta/db"
	"fmt"
)

type UsersReadOnlyRepo struct {
	db db.DB
}

func NewUsersRepo(dbFilePath string) *UsersReadOnlyRepo {
	return &UsersReadOnlyRepo{
		db: *db.NewDB(dbFilePath),
	}
}

func (UR *UsersReadOnlyRepo) GetById(id string) (entity.User, error) {
	users, err := UR.GetUsers()

	if err != nil {
		return nil, err
	}

	for _, value := range users {
		if id == value.GetId() {
			return value, nil

		}
	}

	return nil, fmt.Errorf("User with id: %s not found", id)
}

func (UR *UsersReadOnlyRepo) GetIdByEmail(email string) (string, error) {
	users, err := UR.GetUsers()

	if err != nil {
		return "nil", err
	}

	for _, value := range users {
		if email == value.GetEmail() {
			return value.GetId(), nil
		}
	}

	return "", fmt.Errorf("User with email: %s not found", email)
}

func (UR *UsersReadOnlyRepo) GetUsers() ([]entity.User, error) {

	wapperArray, err := UR.db.GetDB()

	if err != nil {
		return nil, err
	}

	var userSlice []entity.User = make([]entity.User, len(wapperArray))

	for idx, value := range wapperArray {
		userSlice[idx] = value.User
	}

	return userSlice, nil
}
