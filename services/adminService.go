package services

import (
	"fmt"
	"log/slog"

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

func (AS *AdminService) UpdateEmployeeDetails(email, newData, field string) error {
	session, ok := global.GetGlobalSession()

	if !ok {
		return notLoggedInError
	}

	switch field {
	case "department":
		if session.UserRole != "admin" {
			slog.Warn("ADMIN SERIVCE: Role is not admin for user", "email", session.UserEmail)
			return forbiddenError
		}
		return AS.employeeRepo.UpdateEmployee(email, newData, field)
	case "position":
		if session.UserRole != "admin" {
			slog.Warn("ADMIN SERIVCE: Role is not admin for user", "email", session.UserEmail)
			return forbiddenError
		}
		err := AS.employeeRepo.UpdateEmployee(email, newData, field)

		if err == nil {
			slog.Info("EMPLOYEE SERIVCE: Employee data updated", "effected", email, "field", field ,"new", newData, "principal", session.UserEmail)
		}
		return err
	default:
		slog.Error("ADMIN SERIVCE: Invalemail field", "field", field, "email", session.UserEmail)
		return fmt.Errorf("Invalemail Field %s, Vaild are position, department", field)
	}
}

func (AS *AdminService) DeleteEmployee(email string) error {
	session, ok := global.GetGlobalSession()

	if !ok {
		return notLoggedInError
	}

	if session.UserRole != "admin" {
		slog.Warn("ADMIN SERIVCE: Role is not admin for user", "email", session.UserEmail)
		return forbiddenError
	}

	err := AS.employeeRepo.DeleteEmployee(email)

	if err == nil {
		slog.Info("ADMIN SERIVCE: Employee deleted", "effected", email, "principal", session.UserEmail)
	}

	return err
}

func (AS *AdminService) GetAllEmployees() ([]entity.Employee, error) {
	session, ok := global.GetGlobalSession()

	if !ok {
		return nil, notLoggedInError
	}

	if session.UserRole != "admin" {
		slog.Warn("ADMIN SERIVCE: Role is not admin for user", "email", session.UserEmail)
		return nil, forbiddenError
	}

	employees, err := AS.employeeRepo.GetAllEmployees()

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (AS *AdminService) AdminExists() (bool, error) {
	admins, err := AS.adminRepo.GetAllAdmins()
	if err != nil {
		return false, nil
	}

	return admins != nil, nil
}