package services

import (
	"fmt"

	"wg.dde1ta/entity"
	"wg.dde1ta/global"
	"wg.dde1ta/repo"
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

	users, err := AS.usersRepo.GetUsers()

	if err != nil {
		return nil, err
	}

	var employees []entity.Employee;

	for _, value := range users{
		if value.GetRole() == "employee" {
			employees = append(employees, value.(entity.Employee))
		}
	}

	return employees, nil
}