package user

var singleUser User

type User interface {
	CheckLogin(string,string) bool
}

type user struct {
	id string
	password string
}

func NewUser(id, password string) User{
	return &user{id:id,password:password}
}

func (u *user)CheckLogin(id,password string) bool{
	return u.checkId(id) && u.checkPassword(password)
}

func (u *user)checkId(id string) bool {
	return u.id == id
}

func (u *user)checkPassword(password string) bool {
	return u.password == password
}

func GetSingleUser() User{
	if singleUser == nil {
		singleUser = NewUser("jack","1234")
	}

	return singleUser
}