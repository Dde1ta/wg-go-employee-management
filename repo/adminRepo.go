package repo

import (
	"encoding/json"
	"fmt"
	"wg.dde1ta/db"
	"wg.dde1ta/entity"
)

type AdminRepo struct {
	db db.DB
}

func NewAdminRepo(dbFilePath string) *AdminRepo {
	return &AdminRepo{
		db: *db.NewDB(dbFilePath),
	}
}

func (AR *AdminRepo) getUsers() ([]entity.User, error) {

	wapperArray, err := AR.db.GetDB()

	if err != nil {
		return nil, err
	}

	var userSlice []entity.User = make([]entity.User, len(wapperArray))

	for idx, value := range wapperArray {
		userSlice[idx] = value.User
	}

	return userSlice, nil
}

func (AR *AdminRepo) saveToDB(array []entity.User) error {
	toSave, err := json.Marshal(array)

	if err != nil {
		return err
	}

	err = AR.db.SaveToDB(string(toSave))

	return err
}

func (ER *AdminRepo) CreateAdmin(email, password string) (string, error) {
	newAdminObj, err := entity.NewAdmin(
		email, password,
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

	users = append(users, newAdminObj)

	err = ER.saveToDB(users)

	if err != nil {
		return "", err
	}

	return newAdminObj.Id, nil
}

func (ER *AdminRepo) GetById(id string) (*entity.Admin, error) {
	users, err := ER.getUsers()

	if err != nil {
		return nil, err
	}

	for _, value := range users {
		if id == value.GetId() {
			if emp, ok := value.(entity.Admin); ok { 
				return &emp, nil
			 }
		}
	}

	return nil, fmt.Errorf("Admin with id: %s not found", id)
}

func (ER *AdminRepo) GetIdByEmail(email string) (string, error) {
	users, err := ER.getUsers()

	if err != nil {
		return "nil", err
	}

	for _, value := range users {
		if email == value.GetEmail() {
			if emp, ok := value.(entity.Admin); ok { 
				return emp.Id, nil
			 }
		}
	}

	return "", fmt.Errorf("Admin with email: %s not found", email)
}

func (ER *AdminRepo) UpdateAdmin(id string, newData string, field string) error {
	/**
	Valid Fields := Name, Email, Password, Contact Number
	*/

	users, err := ER.getUsers()

	if err != nil {
		return err
	}

	var AdminToUpdate *entity.Admin = nil
	var AdminIndex int = 0;

	for idx, value := range users {
		
		if id == value.GetId() && value.GetRole() == "Admin" {
			AdminIndex = idx
			emp, ok := value.(entity.Admin)
			if ok {
				AdminToUpdate = &emp
			}
			break
		}
	}

	if AdminToUpdate == nil {
		return fmt.Errorf("Admin with Id %s does not exist", id)
	}

	switch field {
	case "email":
		AdminToUpdate.Email = newData
	case "password":
		AdminToUpdate.PasswordHashed = newData
	default:
		return fmt.Errorf("Invalid Field %s Valid fields are email, password", field)
	}
	if AdminToUpdate.Validate() != nil {
		return fmt.Errorf("Incorrect Format for field %s, value %s is invalid", field, newData)
	}

	users[AdminIndex] = *AdminToUpdate

	return ER.saveToDB(users)
}
