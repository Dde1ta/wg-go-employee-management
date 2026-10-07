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


func (UR *UsersReadOnlyRepo) GetUserByEmail(email string) (*entity.User, error) {
	users, err := UR.getUsers()

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userEmail, _ := user.GetProperty("email")
		if email == userEmail.ToString(){
			return &user, nil
		}
	}

	return nil, fmt.Errorf("User with email: %s not found", email)
}

func (UR *UsersReadOnlyRepo) IsUniqueEmail(email string) (bool, error) {
	users, err := UR.getUsers()

	if err != nil {
		return false, err
	}

	for _, user := range users {
		userEmail, _ := user.GetProperty("email")
		if email == userEmail.ToString(){
			return false, nil
		}
	}

	return true, nil
}

func (UR *UsersReadOnlyRepo) IsUniqueContact(contact string) (bool, error) {
	users, err := UR.getUsers()

	if err != nil {
		return false, err
	}

	for _, user := range users {
		userContact, err := user.GetProperty("contact")
		if err == nil {
			if contact == userContact.ToString(){
				return false, nil
			}
		}
	}

	return true, nil
}
