package services

import (
	"errors"
	"fmt"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
	"wg.dde1ta/global"
	"wg.dde1ta/repo"
)

var notLoggedInError error = errors.New("You are not logged in / Invalid Session")
var forbiddenError error = errors.New("Forbidden action")

type AuthService struct {
	employeeRepo *repo.EmployeeRepo
	usersRepo    *repo.UsersReadOnlyRepo
	adminRepo    *repo.AdminRepo
}

func NewAuthService(dbFilePath string) AuthService {
	return AuthService{
		employeeRepo: repo.NewEmployeeRepo(dbFilePath),
		adminRepo:    repo.NewAdminRepo(dbFilePath),
		usersRepo:    repo.NewUsersRepo(dbFilePath),
	}
}

func (auth *AuthService) CreateAdmin(email, password string) error {
	session, ok := global.GetGlobalSession()

	if !ok {
		slog.Warn("AUTH SERIVCE: system not in setup Mode")
		return fmt.Errorf("Not in setup mode")
	}

	if session.UserRole != "setup" {
		slog.Warn("AUTH SERIVCE: Role is not Setup for user", "email", session.UserEmail)
		return fmt.Errorf("Forbidden")
	}

	err := auth.adminRepo.CreateAdmin(email, password)

	if err != nil {
		return err
	}

	slog.Info("AUTH SERIVCE: New admin created", "newAdmin", email, "principal", session.UserEmail)
	return nil
}

func (auth *AuthService) SignUpEmployee(name, email, password, phone, department, position string) error {
	return auth.employeeRepo.CreateEmployee(name, email, password, phone, department, position)
}

func (auth *AuthService) Login(email, givenPassword string) error {
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

	slog.Info("AUTH SERIVCE: Login by", "userEmail", email)

	return nil
}
