package models

type User struct {
	Userid   uint `gorm:"primarykey"`
	Username string
	Password string
	Profile  Profile `gorm:"foreignkey:Profileid"`
}

type Profile struct {
	Userid    uint `gorm:"primarykey"`
	Profileid uint
	Name      string
	Age       int
}
