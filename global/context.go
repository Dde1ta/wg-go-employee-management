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
	globalSessionContext = SessionContext{
		UserId: user.GetId(),
		UserRole: user.GetRole(),
		UserEmail: user.GetEmail(),
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
