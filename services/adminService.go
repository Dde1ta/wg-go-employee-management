package services

import (
	"fmt"
	"wg.dde1ta/repo"
	"wg.dde1ta/entity"
	"wg.dde1ta/global"
)

type AdminService struct {
	employeeRepo *repo.EmployeeRepo
	adminRepo    *repo.AdminRepo
	usersRepo    *repo.UsersReadOnlyRepo
}

func NewAdminService(dbFilePath string) AdminService{
	return AdminService{
		employeeRepo: repo.NewEmployeeRepo(dbFilePath),
		adminRepo: repo.NewAdminRepo(dbFilePath),
		usersRepo: repo.NewUsersRepo(dbFilePath),
	}
}

func (AS *AdminService) UpdateEmployeeDetails(id, newData, field string) error {
	session, ok := global.GetGlobalSession()

	if !ok {
		return notLoggedInError
	}

	switch field {
	case "department":
		if session.UserRole != "admin" {
			return forbiddenError
		}
		return AS.employeeRepo.UpdateEmployee(id, newData, field)
	case "position":
		if session.UserRole != "admin" {
			return forbiddenError
		}
		return AS.employeeRepo.UpdateEmployee(id, newData, field)
	default:
		return fmt.Errorf("Invalid Field %s, Vaild are position, department", field)
	}
}

func (AS *AdminService) DeleteEmployee(id string) error {
	session, ok := global.GetGlobalSession()

	if !ok {
		return notLoggedInError
	}

	if session.UserRole != "admin" {
		return forbiddenError
	}

	return AS.employeeRepo.DeleteEmployee(id)
}

func (AS *AdminService) GetAllEmployees() ([]entity.Employee, error) {
	session, ok := global.GetGlobalSession()

	if !ok {
		return nil, notLoggedInError
	}

	if session.UserRole != "admin" {
		return nil, forbiddenError
	}

	employees, err := AS.employeeRepo.GetAllEmployees()

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (auth *AuthService) AdminExists() (bool, error) {
	admins, err := auth.adminRepo.GetAllAdmins()
	if err != nil {
		return false, nil
	}

	return admins != nil, nil
}