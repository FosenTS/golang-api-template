package dto

type UserCreate struct {
	Login      string
	Password   string
	Permission uint
}
