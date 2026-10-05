package entity

type User interface{
	ToJson() (string, error)
	GetId() (string)
	GetRole() (string)
	GetEmail() (string)
	GetPassword() (string)
	Validate() (error)
}