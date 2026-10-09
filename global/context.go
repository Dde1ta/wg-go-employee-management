package global

import (
	"wg.dde1ta/entity"
)

var globalSessionContext SessionContext

type SessionContext struct {
	UserRole  string
	UserEmail string
}

func NewSessionContext(user entity.User) SessionContext {
	roleS, _ := user.GetProperty("role")
	emailS, _ := user.GetProperty("email")

	globalSessionContext = SessionContext{
		UserRole:  roleS.ToString(),
		UserEmail: emailS.ToString(),
	}

	return globalSessionContext
}

func GetGlobalSession() (SessionContext, bool) {
	if globalSessionContext == (SessionContext{}) {
		return SessionContext{}, false
	}
	return globalSessionContext, true
}

func SetSetupRole() {
	globalSessionContext.UserRole = "setup"
}

func UpdateSessionEmail(email string) {
	globalSessionContext.UserEmail = email
}

func LogOut() {
	globalSessionContext = SessionContext{}
}
