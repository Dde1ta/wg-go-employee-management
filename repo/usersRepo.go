package repo


import (
	"wg.dde1ta/entity"
	"wg.dde1ta/db"
	"fmt"
)

type UsersReadOnlyRepo struct {
	repo
}

func NewUsersRepo(dbFilePath string) *UsersReadOnlyRepo {
	return &UsersReadOnlyRepo{
		db: *db.NewDB(dbFilePath),
	}
}

func (UR *UsersReadOnlyRepo) GetById(id string) (entity.User, error) {
	users, err := UR.getUsers()

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
	users, err := UR.getUsers()

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

func (UR *UsersReadOnlyRepo) IsUniqueEmail(email string) (bool, error) {
	users, err := UR.getUsers()

	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user.GetEmail() == email{
			return false, nil
		}
	}

	return true, nil
}

func (UR *UsersReadOnlyRepo) IsUniquePhone(phone string) (bool, error) {
	users, err := UR.getUsers()

	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user.GetEmail() == phone{
			return false, nil
		}
	}

	return true, nil
}
