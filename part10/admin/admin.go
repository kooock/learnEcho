package admin




type Admin interface {
	CheckAdminInfo(string,string) bool
}

type admin struct {
	id string
	password string
}

func New(id string,password string) Admin{
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
