package repo


import (
	"wg.dde1ta/entity"
	"fmt"
)

type UsersReadOnlyRepo struct {
	repo
}

func NewUsersRepo(dbFilePath string) *UsersReadOnlyRepo {
	return &UsersReadOnlyRepo{
		repo: *NewRepo(dbFilePath),
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
	fmt.Println("Debug: Checking Unique Update")

	users, err := UR.getUsers()

	if err != nil {
		return false, err
	}

	for _, user := range users {
		userContact, err := user.GetProperty("contact")

		if err == nil && userContact != nil {
			if contact == userContact.ToString(){
				return false, nil
			}
		}
	}

	return true, nil
}
