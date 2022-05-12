package model


//记录课程信息
type Course struct {
	Id uint
	CourseId string `json:"course_id" form:"course_id"`
	Name string `json:"name" form:"name" binding:"required"`
	Weekday uint `json:"weekday" form:"weekday" binding:"required"`
	Class     uint `json:"class" form:"class" binding:"required"`
	Max uint `json:"max" form:"max" binding:"required"`
	Now uint `json:"now" form:"now"`
	TeacherId string `json:"teacher_id" form:"teacher_id" binding:"required"`
}
