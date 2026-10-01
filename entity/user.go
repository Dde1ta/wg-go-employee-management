package entity

type User interface{
	ToJson() (string, error)
	GetId() (string)
}