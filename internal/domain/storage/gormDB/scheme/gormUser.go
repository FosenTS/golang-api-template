package scheme

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Login      string `gorm:"index;unique"`
	Password   string
	Permission uint
}
