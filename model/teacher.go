package model

type Teacher struct {
	Id uint
	TeacherId string `json:"teacher_id" form:"teacher_id"`
	Name  string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email string `json:"email" form:"email" binding:"required" gorm:"unique" validate:"email"`
}
