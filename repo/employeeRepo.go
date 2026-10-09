package repo

import (
	"fmt"
	"log/slog"
	"runtime/debug"
	"slices"

	"wg.dde1ta/entity"
)

type EmployeeRepo struct {
	repo
	UsersReadOnlyRepo
}

func NewEmployeeRepo(dbFilePath string) *EmployeeRepo {
	return &EmployeeRepo{
		repo:              *NewRepo(dbFilePath),
		UsersReadOnlyRepo: *NewUsersRepo(dbFilePath),
	}
}

func (ER *EmployeeRepo) CreateEmployee(name, email, password, phone, department, position string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// 2. Log the panic value and the stack trace as slog attributes
			slog.Error("recovered from panic",
				slog.Any("panic_value", r),
				slog.String("stack", string(debug.Stack())),
			)

			err = fmt.Errorf("Unexcepted err see logs")
		}
	}()

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

	isUnique, err = ER.IsUniqueContact(phone)

	if err != nil {
		return err
	}

	if !isUnique {
		return fmt.Errorf("%s Phone is already in use", phone)
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

	slog.Info("EMPLOYEE REPO: New Employee created: ", "email", email, "name", name)

	return nil
}

func (ER *EmployeeRepo) GetEmployeeByEmail(email string) (emp *entity.Employee, err error) {
	defer func() {
		if r := recover(); r != nil {
			// 2. Log the panic value and the stack trace as slog attributes
			slog.Error("recovered from panic",
				slog.Any("panic_value", r),
				slog.String("stack", string(debug.Stack())),
			)
			emp = nil
			err = fmt.Errorf("Unexcepted error see logs")
		}
	}()

	users, err := ER.getUsers()

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userEmail, _ := user.GetProperty("email")
		userRole, _ := user.GetProperty("role")
		if email == userEmail.ToString() && userRole.ToString() == "employee" {
			if emp, ok := user.(*entity.Employee); ok {
				return emp, nil
			}
		}
	}

	return nil, fmt.Errorf("Employee with email: %s not found", email)
}

func (ER *EmployeeRepo) UpdateEmployee(email string, newData string, field string) (err error) {
	/**
	Valid Fields := Name, Email, Password, Contact Number
	*/

	defer func() {
		if r := recover(); r != nil {
			// 2. Log the panic value and the stack trace as slog attributes
			slog.Error("recovered from panic",
				slog.Any("panic_value", r),
				slog.String("stack", string(debug.Stack())),
			)

			err = fmt.Errorf("Unexcepted err see logs")
		}
	}()

	users, err := ER.getUsers()

	if err != nil {
		return err
	}

	var employeeToUpdate *entity.Employee = nil
	var employeeIndex int = 0

	for idx, user := range users {
		userEmail, _ := user.GetProperty("email")
		userRole, _ := user.GetProperty("role")
		if email == userEmail.ToString() && userRole.ToString() == "employee" {
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

	var value entity.MyString
	value.CopyString(newData)

	err = employeeToUpdate.SetProperty(field, &value)

	if err != nil {
		return err
	}

	users[employeeIndex] = employeeToUpdate

	return ER.saveToDB(users)
}

func (ER *EmployeeRepo) DeleteEmployee(email string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// 2. Log the panic value and the stack trace as slog attributes
			slog.Error("recovered from panic",
				slog.Any("panic_value", r),
				slog.String("stack", string(debug.Stack())),
			)

			err = fmt.Errorf("Unexcepted err see logs")
		}
	}()

	users, err := ER.getUsers()

	if err != nil {
		return err
	}

	var foundAndDeleted bool = false

	for idx, user := range users {
		userEmail, _ := user.GetProperty("email")
		userRole, _ := user.GetProperty("role")
		if userEmail.ToString() == email && userRole.ToString() == "employee" {
			users = slices.Delete(users, idx, idx+1)
			foundAndDeleted = true
			break
		}
	}

	if !foundAndDeleted {
		return fmt.Errorf("Employe with email %s not found", email)
	}

	return ER.saveToDB(users)
}

func (ER *EmployeeRepo) GetAllEmployees() (employeeList []entity.Employee, err error) {
	defer func() {
		if r := recover(); r != nil {
			// 2. Log the panic value and the stack trace as slog attributes
			slog.Error("recovered from panic",
				slog.Any("panic_value", r),
				slog.String("stack", string(debug.Stack())),
			)
			employeeList = nil
			err = fmt.Errorf("Unexcepted err see logs")
		}
	}()

	users, err := ER.getUsers()

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userRole, _ := user.GetProperty("role")
		if userRole.ToString() == "employee" {
			if emp, ok := user.(*entity.Employee); ok {
				employeeList = append(employeeList, *emp)
			}
		}
	}

	return employeeList, nil
}
