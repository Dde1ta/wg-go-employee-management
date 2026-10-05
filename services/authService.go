package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"wg.dde1ta/global"
	"wg.dde1ta/repo"
)

type AuthService struct {
	employeeRepo *repo.EmployeeRepo
	usersRepo *repo.UsersReadOnlyRepo
	adminRepo *repo.AdminRepo
}

func NewAuthService(dbFile string) *AuthService {
	return &AuthService{
		employeeRepo: repo.NewEmployeeRepo(dbFile),
		adminRepo: repo.NewAdminRepo(dbFile),
	}
}

func (auth *AuthService) CreateAdmin(email, password string) (string, error) {
	return auth.adminRepo.CreateAdmin(email, password)
}

func (auth *AuthService) SignUpEmployee(name, email, password, phone, department, position string) (string, error) {
	return auth.employeeRepo.CreateEmployee(name, email, password, phone, department, position)
}

func (auth *AuthService) Login(email, givenPassword string) (error) {
	id, err := auth.usersRepo.GetIdByEmail(email)

	if err != nil {
		return err
	}

	user, err := auth.usersRepo.GetById(id)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(givenPassword))

	if err != nil {
		return errors.New("The Password or Email is invalid")
	}

	global.NewSessionContext(user)

	return nil
}
