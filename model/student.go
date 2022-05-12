package model

type Student struct {
	Id uint
	StudentId string `json:"student_id" form:"student_id"`
	Name  string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email string `json:"email" form:"email" binding:"required" gorm:"unique" validate:"email"`
}