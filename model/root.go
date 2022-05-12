package model

type Root struct {
	Id uint `gorm:"auto_increment=100"`
	Name  string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}
