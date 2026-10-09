package repo

import (
	"fmt"
	"log/slog"
	"runtime/debug"

	"wg.dde1ta/entity"
)

type AdminRepo struct {
	repo
	UsersReadOnlyRepo
}

func NewAdminRepo(dbFilePath string) *AdminRepo {
	return &AdminRepo{
		repo:              *NewRepo(dbFilePath),
		UsersReadOnlyRepo: *NewUsersRepo(dbFilePath),
	}
}

func (AR *AdminRepo) CreateAdmin(email, password string) (err error) {

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

	newAdminObj, err := entity.NewAdmin(
		email, password,
	)

	if err != nil {
		return err
	}

	isUnique, err := AR.IsUniqueEmail(email)

	if err != nil {
		return err
	}

	if !isUnique {
		return fmt.Errorf("%s email is already in use", email)
	}

	users, err := AR.getUsers()

	if err != nil {
		return err
	}

	users = append(users, newAdminObj)

	err = AR.saveToDB(users)

	if err != nil {
		return err
	}

	return nil
}

func (AR *AdminRepo) GetAdminByEmail(email string) (adm *entity.Admin, err error) {
	defer func() {
		if r := recover(); r != nil {
			// 2. Log the panic value and the stack trace as slog attributes
			slog.Error("recovered from panic",
				slog.Any("panic_value", r),
				slog.String("stack", string(debug.Stack())),
			)

			adm = nil
			err = fmt.Errorf("Unexcepted err see logs")
		}
	}()

	users, err := AR.getUsers()

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userEmail, _ := user.GetProperty("email")
		userRole, _ := user.GetProperty("role")
		if email == userEmail.ToString() && userRole.ToString() == "admin" {
			if adm, ok := user.(*entity.Admin); ok {
				return adm, nil
			}
		}
	}

	return nil, fmt.Errorf("Admin with email: %s not found", email)
}

func (AR *AdminRepo) GetAllAdmins() (adminList []entity.Admin, err error) {
	defer func() {
		if r := recover(); r != nil {
			// 2. Log the panic value and the stack trace as slog attributes
			slog.Error("ADMIN REPO: recovered from panic",
				slog.Any("panic_value", r),
				slog.String("stack", string(debug.Stack())),
			)

			adminList = nil
			err = fmt.Errorf("Unexcepted err see logs")
		}
	}()

	users, err := AR.getUsers()

	if err != nil {
		return nil, err
	}

	adminList = nil

	for _, user := range users {
		userRole, _ := user.GetProperty("role")
		if userRole.ToString() == "admin" {
			if adm, ok := user.(*entity.Admin); ok {
				adminList = append(adminList, *adm)
			}
		}
	}

	return adminList, nil
}
