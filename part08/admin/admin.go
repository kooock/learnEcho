package admin

import (
	"github.com/labstack/echo"
)

type Admin interface {
	CheckAdminInfo(string,string) bool
	SetAdminGroup(g *echo.Group)
}

type admin struct {
	id string
	password string
}

func NewAdmin(id string,password string) Admin{
	return &admin{id:id,password:password}
}

func (ad *admin)CheckAdminInfo(id string,password string) bool{

	return ad.checkId(id) && ad.checkPassword(password)
}

func (ad *admin)checkId(id string) bool{
	return ad.id == id
}

func (ad *admin)checkPassword(password string) bool{
	return ad.password == password
}

func (ad *admin)SetAdminGroup(g *echo.Group) {
	setLogger(g)
	setAuth(g,ad)
}