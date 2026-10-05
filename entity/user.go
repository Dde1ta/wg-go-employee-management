package entity

type User interface{
	ToJson() (string, error)
	GetId() (string)
	GetRole() (string)
	GetEmail() (string)
	Validate() (error)
}