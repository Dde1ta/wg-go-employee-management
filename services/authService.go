package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"wg.dde1ta/global"
	"wg.dde1ta/repo"
)

var notLoggedInError error = errors.New("You are not logged in / Invalid Session")
var forbiddenError error = errors.New("Forbidden action")

type AuthService struct {
	employeeRepo *repo.EmployeeRepo
	usersRepo *repo.UsersReadOnlyRepo
	adminRepo *repo.AdminRepo
}

func NewAuthService(dbFilePath string) AuthService {
	return AuthService{
		employeeRepo: repo.NewEmployeeRepo(dbFilePath),
		adminRepo: repo.NewAdminRepo(dbFilePath),
		usersRepo: repo.NewUsersRepo(dbFilePath),
	}
}

func (auth *AuthService) CreateAdmin(email, password string) (error) {
	return auth.adminRepo.CreateAdmin(email, password)
}

func (auth *AuthService) SignUpEmployee(name, email, password, phone, department, position string) (error) {
	return auth.employeeRepo.CreateEmployee(name, email, password, phone, department, position)
}

func (auth *AuthService) Login(email, givenPassword string) (error) {
	user, err := auth.usersRepo.GetUserByEmail(email)

	if err != nil {
		return err
	}

	userCopy := *user

	pass, _ := userCopy.GetProperty("password")

	err = bcrypt.CompareHashAndPassword([]byte(pass.ToString()), []byte(givenPassword))

	if err != nil {
		return errors.New("The Password or Email is invalid")
	}

	global.NewSessionContext(userCopy)

	return nil
}
