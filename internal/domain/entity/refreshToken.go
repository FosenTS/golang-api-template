package entity

type RefreshToken struct {
	ID                 uint
	Token              string
	Login              string
	ExpirationTimeUnix int
	CreateTimeUnix     int
}
