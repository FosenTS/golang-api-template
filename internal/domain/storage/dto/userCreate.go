package dto

type UserCreate struct {
	Login      string `json:"login" binding:"true"`
	Password   string `json:"password" binding:"true"`
	Permission uint   `json:"permission" binding:"true"`
}
