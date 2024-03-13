package safeobject

type User struct {
	Login      string `json:"login"`
	Permission uint   `json:"permission"`
}

func NewUser(login string, permission uint) *User {
	return &User{Login: login, Permission: permission}
}
