package dto

type RefreshTokenCreate struct {
	Token              string
	Login              string
	ExpirationTimeUnix int
	CreateTimeUnix     int
}
