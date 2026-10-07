package repo

import (
	"fmt"
	"wg.dde1ta/entity"
)

type AdminRepo struct {
	repo
	UsersReadOnlyRepo
}

func NewAdminRepo(dbFilePath string) *AdminRepo {
	return &AdminRepo{
		repo: *NewRepo(dbFilePath),
		UsersReadOnlyRepo: *NewUsersRepo(dbFilePath),
	}
}

func (AR *AdminRepo) CreateAdmin(email, password string) (error) {
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

func (AR *AdminRepo) GetAdminByEmail(email string) (*entity.Admin, error) {
	users, err := AR.getUsers()

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userEmail, _ := user.GetProperty("email")
		userRole, _ := user.GetProperty("role")
		if email == userEmail.ToString() && userRole.ToString() == "admin"{
			if adm, ok := user.(*entity.Admin); ok { 
				return adm, nil
			 }
		}
	}

	return nil, fmt.Errorf("Admin with email: %s not found", email)
}

func (AR *AdminRepo) GetAllAdmins() ([]entity.Admin, error) {
	users, err := AR.getUsers()
	
	if err != nil {
		return nil, err
	}

	var adminList []entity.Admin = nil

	for _, user := range users {
		userRole, _ := user.GetProperty("role")
		if userRole.ToString() == "admin" {
			if adm, ok := user.(*entity.Admin); ok{
				adminList = append(adminList, *adm)
			}
		}
	}

	return adminList, nil
}
