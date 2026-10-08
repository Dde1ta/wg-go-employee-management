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

func (ES *EmployeeService) GetEmployeeByEmail(email string) (*entity.Employee, error){
	session, ok := global.GetGlobalSession()

	if !ok {
		return nil, errors.New("You are not logged in / Invalid Session")
	}

	if session.UserRole == "admin" || email == session.UserEmail {
		return ES.employeeRepo.GetEmployeeByEmail(email)
	}

	return nil, errors.New("Forbidden action")
}

func (ES *EmployeeService) UpdateEmployeeDetails(email, newData, field string) (error) {
	fmt.Printf("DEBUG: Update call %s %s %s\n", email, newData, field)
	session, ok := global.GetGlobalSession()

	if !ok {
		return notLoggedInError
	}

	if field == "department" || field == "position" {
		if session.UserRole != "admin" {
			return forbiddenError
		}
		return ES.employeeRepo.UpdateEmployee(email, newData, field)
	}

	if session.UserEmail != email{
		return forbiddenError
	}
	err := ES.employeeRepo.UpdateEmployee(email, newData, field)

	if err != nil {
		return err
	}

	if field == "email"{
		global.UpdateSessionEmail(newData)
	}

	return nil
}
