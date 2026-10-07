package global

import (
	"wg.dde1ta/entity"
)

var globalSessionContext SessionContext;

type SessionContext struct {
	UserId string
	UserRole string
	UserEmail string
}

func NewSessionContext(user entity.User) SessionContext{
	idS, _ := user.GetProperty("id")
	roleS, _ := user.GetProperty("role")
	emailS, _ := user.GetProperty("email")

	globalSessionContext = SessionContext{
		UserId: idS.ToString(),
		UserRole: roleS.ToString(),
		UserEmail: emailS.ToString(),
	}

	return globalSessionContext
}

func GetGlobalSession() (SessionContext, bool) {
	if globalSessionContext == (SessionContext{}){
		return SessionContext{}, false
	}
	return globalSessionContext, true
}

func LogOut() {
	globalSessionContext = SessionContext{}
}
