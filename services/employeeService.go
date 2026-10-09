package services

import (
	"errors"
	"log/slog"

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

	slog.Warn("EMPLOYEE SERIVCE: forbidden request to get employee", "email=", session.UserEmail)
	return nil, errors.New("Forbidden action")
}

func (ES *EmployeeService) UpdateEmployeeDetails(email, newData, field string) (error) {
	session, ok := global.GetGlobalSession()

	if !ok {
		return notLoggedInError
	}

	if field == "department" || field == "position" {
		if session.UserRole != "admin" {
			slog.Warn("EMPLOYEE SERIVCE: forbidden request to update (Not Admin)", "affected=", email, "field=", field, "principal=", session.UserEmail)
			return forbiddenError
		}
		return ES.employeeRepo.UpdateEmployee(email, newData, field)
	}

	if session.UserEmail != email{
		slog.Warn("EMPLOYEE SERIVCE: forbidden request to update", "affected=", email, "field=", field, "principal=", session.UserEmail)
		return forbiddenError
	}
	err := ES.employeeRepo.UpdateEmployee(email, newData, field)

	if err != nil {
		return err
	}

	if field == "email"{
		global.UpdateSessionEmail(newData)
		slog.Info("EMPLOYEE SERIVCE: Session updated", "previous=", email, "new=", newData)
	}
	slog.Info("EMPLOYEE SERIVCE: Employee data updated", "effected=", email, "field=", field ,"new=", newData, "principal=", session.UserEmail)

	return nil
}
