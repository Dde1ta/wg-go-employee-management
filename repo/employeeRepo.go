package repo

import (
	"encoding/json"
	"fmt"
	"slices"
	"wg.dde1ta/db"
	"wg.dde1ta/entity"
)

type EmployeeRepo struct {
	db db.DB
}

func NewEmployeeRepo(dbFilePath string) *EmployeeRepo {
	return &EmployeeRepo{
		db: *db.NewDB(dbFilePath),
	}
}

func (ER *EmployeeRepo) getUsers() ([]entity.User, error) {

	wapperArray, err := ER.db.GetDB()

	if err != nil {
		return nil, err
	}

	var userSlice []entity.User = make([]entity.User, len(wapperArray))

	for idx, value := range wapperArray {
		userSlice[idx] = value.User
	}

	return userSlice, nil
}

func (ER *EmployeeRepo) CreateEmployee(name, email, password, phone, department, position string) (string, error) {
	newEmployeeObj, err := entity.NewEmployee(
		name, email, password, phone, department, position,
	)

	if err != nil {
		return "", err
	}

	users, err := ER.getUsers()

	if err != nil {
		return "", err
	}

	for _, values := range users{
		if values.GetEmail() == email {
			return "", fmt.Errorf("Email %s is not unique", email)
		}
	}

	users = append(users, newEmployeeObj)

	err = ER.saveToDB(users)

	if err != nil {
		return "", err
	}

	return newEmployeeObj.Id, nil
}

func (ER *EmployeeRepo) UpdateEmployee(id string, newData string, field string) error {
	/**
	Valid Fields := Name, Email, Password, Contact Number
	*/

	users, err := ER.getUsers()

	if err != nil {
		return err
	}

	var employeeToUpdate *entity.Employee = nil
	var employeeIndex int = 0;

	for idx, value := range users {
		
		if id == value.GetId() && value.GetRole() == "employee" {
			employeeIndex = idx
			emp, ok := value.(entity.Employee)
			if ok {
				employeeToUpdate = &emp
			}
			break
		}
	}

	if employeeToUpdate == nil {
		return fmt.Errorf("Employee with Id %s does not exist", id)
	}

	switch field {
	case "name":
		employeeToUpdate.Name = newData
	case "email":
		employeeToUpdate.Email = newData
	case "contact_number":
		employeeToUpdate.Phone = newData
	case "password":
		employeeToUpdate.PasswordHashed = newData
	case "position":
		employeeToUpdate.Position = newData
	case "department":
		employeeToUpdate.Department = newData
	default:
		return fmt.Errorf("Invalid Field %s Valid fields are name, email, contact_number, password", field)
	}
	if employeeToUpdate.Validate() != nil {
		return fmt.Errorf("Incorrect Format for field %s, value %s is invalid", field, newData)
	}

	users[employeeIndex] = *employeeToUpdate

	return ER.saveToDB(users)
}

func (ER *EmployeeRepo) GetById(id string) (*entity.Employee, error) {
	users, err := ER.getUsers()

	if err != nil {
		return nil, err
	}

	for _, value := range users {
		if id == value.GetId() {
			if emp, ok := value.(entity.Employee); ok { 
				return &emp, nil
			 }
		}
	}

	return nil, fmt.Errorf("Employee with id: %s not found", id)
}

func (ER *EmployeeRepo) DeleteEmployee(id string) error {
	users, err := ER.getUsers()

	if err != nil {
		return err
	}

	for idx, value := range users {
		if id == value.GetId() && value.GetRole() == "employee"{
			users = slices.Delete(users, idx, idx + 1)
			break
		}
	}

	return ER.saveToDB(users)
}

func (ER *EmployeeRepo) GetIdByEmail(email string) (string, error) {
	users, err := ER.getUsers()

	if err != nil {
		return "nil", err
	}

	for _, value := range users {
		if email == value.GetEmail() {
			if emp, ok := value.(entity.Employee); ok { 
				return emp.Id, nil
			 }
		}
	}

	return "", fmt.Errorf("Employee with email: %s not found", email)
}

func (ER *EmployeeRepo) saveToDB(array []entity.User) error {
	fmt.Println("Debug: Saving", array)

	toSave, err := json.Marshal(array)

	if err != nil {
		return err
	}

	err = ER.db.SaveToDB(string(toSave))

	return err
}
