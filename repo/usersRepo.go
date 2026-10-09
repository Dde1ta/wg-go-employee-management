package repo


import (
	"wg.dde1ta/entity"
	"log/slog"
	"runtime/debug"

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


func (UR *UsersReadOnlyRepo) GetUserByEmail(email string) (user *entity.User, err error) {
	defer func() {
		if r := recover(); r != nil {
			// 2. Log the panic value and the stack trace as slog attributes
			slog.Error("recovered from panic",
				slog.Any("panic_value", r),
				slog.String("stack", string(debug.Stack())),
			)

			err = fmt.Errorf("Unexcepted err see logs")
		}
	}()

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

func (UR *UsersReadOnlyRepo) IsUniqueEmail(email string) (ok bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			// 2. Log the panic value and the stack trace as slog attributes
			slog.Error("recovered from panic",
				slog.Any("panic_value", r),
				slog.String("stack", string(debug.Stack())),
			)
			ok = false
			err = fmt.Errorf("Unexcepted err see logs")
		}
	}()
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

func (UR *UsersReadOnlyRepo) IsUniqueContact(contact string) (ok bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			// 2. Log the panic value and the stack trace as slog attributes
			slog.Error("recovered from panic",
				slog.Any("panic_value", r),
				slog.String("stack", string(debug.Stack())),
			)
			ok = false
			err = fmt.Errorf("Unexcepted err see logs")
		}
	}()
	
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
