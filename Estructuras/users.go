package Estructuras

type UserActive struct {
	User     string
	Password string
	Id       string
	Grp      string
	Uid      int
	Gid      int
}

func NewUserActual() UserActive {
	var user UserActive
	user.User = ""
	user.Password = ""
	user.Id = ""
	user.Grp = ""
	user.Uid = -1
	user.Gid = -1
	return user
}

var Logeado UserActive = NewUserActual()
