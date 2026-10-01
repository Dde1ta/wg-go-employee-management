package db

import (
	"wg.dde1ta/entity"
	"encoding/json"
)

type NotFoundEntity struct{
	Id string
}

func NewNotFoundEntity() (NotFoundEntity) {
	newNotFoundEntity := NotFoundEntity{
		Id: "",
	}

	return newNotFoundEntity
}

func (e404 NotFoundEntity) ToJson() (string, error){
	jsonBytes, err := json.Marshal(e404)

	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func (e404 NotFoundEntity) GetId() (string) {
	return e404.Id
}


func GetById(userSlice []entity.User, Id string) (entity.User, error) {
	var s, e int = 0, len(userSlice)
	var mid int = s + (e - s) / 2

	for s <= e {
		midId := userSlice[mid].GetId()

		if midId < Id {
			s = mid + 1
		}

		if midId > Id {
			e = mid - 1
		}

		if midId == Id {
			return userSlice[mid], nil
		}
	} 

	return NewNotFoundEntity(), nil
}
