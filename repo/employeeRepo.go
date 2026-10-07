package repo

import (
	"fmt"
	"slices"
	"wg.dde1ta/entity"
)

type EmployeeRepo struct {
	repo
	UsersReadOnlyRepo
}

func NewEmployeeRepo(dbFilePath string) *EmployeeRepo {
	return &EmployeeRepo{
		repo: *NewRepo(dbFilePath),
		UsersReadOnlyRepo: *NewUsersRepo(dbFilePath),
	}
}

func (ER *EmployeeRepo) CreateEmployee(name, email, password, phone, department, position string) (error) {
	newEmployeeObj, err := entity.NewEmployee(
		name, email, password, phone, department, position,
	)

	if err != nil {
		return err
	}

	isUnique, err := ER.IsUniqueEmail(email)

	if err != nil {
		return err
	}

	if !isUnique {
		return fmt.Errorf("%s email is already in use", email)
	}

	users, err := ER.getUsers()

	if err != nil {
		return err
	}

	users = append(users, newEmployeeObj)

	err = ER.saveToDB(users)

	if err != nil {
		return err
	}

	return nil
}

func (ER *EmployeeRepo) GetEmployeeByEmail(email string) (*entity.Employee, error) {
	users, err := ER.getUsers()

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userEmail, _ := user.GetProperty("email")
		userRole, _ := user.GetProperty("role")
		if email == userEmail.ToString() && userRole.ToString() == "employee"{
			if emp, ok := user.(*entity.Employee); ok { 
				return emp, nil
			 }
		}
	}

	return nil, fmt.Errorf("Employee with email: %s not found", email)
}

func (ER *EmployeeRepo) UpdateEmployee(email string, newData string, field string) error {
	/**
	Valid Fields := Name, Email, Password, Contact Number
	*/
	
	users, err := ER.getUsers()

	if err != nil {
		return err
	}

	var employeeToUpdate *entity.Employee = nil
	var employeeIndex int = 0;


	for idx, user := range users {
		userEmail, _ := user.GetProperty("email")
		userRole, _ := user.GetProperty("role")
		if email == userEmail.ToString() && userRole.ToString() == "employee"{
			if emp, ok := user.(*entity.Employee); ok { 
				employeeToUpdate = emp
				employeeIndex = idx
			 }
		}
	}

	if employeeToUpdate == nil {
		return fmt.Errorf("Employee with email %s Not Found", email)
	}

	switch field {
	case "email":
		ok, err := ER.IsUniqueEmail(newData)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("Email %s is already in use", newData)
		}
	case "contact":
		ok, err := ER.IsUniqueContact(newData)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("Phone %s is already in use", newData)
		}
	}

	var value entity.MyString;
	value.CopyString(newData)

	employeeToUpdate.SetProperty(field, &value)
	
	users[employeeIndex] = employeeToUpdate

	return ER.saveToDB(users)
}

func (ER *EmployeeRepo) DeleteEmployee(email string) error {
	users, err := ER.getUsers()

	if err != nil {
		return err
	}

	for idx, user := range users {
		userEmail, _ := user.GetProperty("email")
		userRole, _ := user.GetProperty("role")
		if userEmail.ToString() == email && userRole.ToString() == "employee"{
			users = slices.Delete(users, idx, idx + 1)
			break
		}
	}

	return ER.saveToDB(users)
}

func (ER *EmployeeRepo) GetAllEmployees() ([]entity.Employee, error) {
	users, err := ER.getUsers()
	
	if err != nil {
		return nil, err
	}

	var employeeList []entity.Employee

	for _, user := range users {
		userRole, _ := user.GetProperty("role")
		if userRole.ToString() == "employee"{
			if emp, ok := user.(*entity.Employee); ok { 
				employeeList = append(employeeList, *emp)
			 }
		}
	}

	return employeeList, nil
} 