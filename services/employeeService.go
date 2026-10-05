package services

import (
	"errors"
	"fmt"
	"wg.dde1ta/entity"
	"wg.dde1ta/global"
	"wg.dde1ta/repo"
)



type EmployeeService struct {
	employeeRepo *repo.EmployeeRepo
}

func NewEmployeeService(dbFilePath string) EmployeeService{
	return EmployeeService{
		employeeRepo: repo.NewEmployeeRepo(dbFilePath),
	}
}

func (ES *EmployeeService) GetEmployeeById(id string) (*entity.Employee, error){
	session, ok := global.GetGlobalSession()

	if !ok {
		return nil, errors.New("You are not logged in / Invalid Session")
	}

	if session.UserRole == "admin" || session.UserId == id {
		return ES.employeeRepo.GetById(id)
	}

	return nil, errors.New("Forbidden action")
}

func (ES *EmployeeService) GetEmployeeIdByEmail(email string) (string, error){
	session, ok := global.GetGlobalSession()

	if !ok {
		return "", notLoggedInError
	}

	if session.UserRole == "admin" || session.UserEmail == email {
		return ES.employeeRepo.GetIdByEmail(email)
	}

	return "", forbiddenError
}

func (ES *EmployeeService) UpdateEmployeeDetails(id, newData, field string) (error) {
	session, ok := global.GetGlobalSession()

	if !ok {
		return notLoggedInError
	}

	switch field{
	case "name":
		if session.UserId != id{
			return forbiddenError
		}
		ES.employeeRepo.UpdateEmployee(id, newData, field)
	case "email":
		if session.UserId != id{
			return forbiddenError
		}
		ES.employeeRepo.UpdateEmployee(id, newData, field)
	case "contact_number":
		if session.UserId != id{
			return forbiddenError
		}
		ES.employeeRepo.UpdateEmployee(id, newData, field)
	case "password":
		if session.UserId != id{
			return forbiddenError
		}
		ES.employeeRepo.UpdateEmployee(id, newData, field)
	default:
		return fmt.Errorf("Invalid Field %s, Vaild are name, email, contact_number, password", field)
	}
	return forbiddenError
}

