package model


//记录学生选课信息
type Record struct {
	Id uint
	Student string `json:"student" form:"student" binding:"required"`
	StudentId string `json:"student_id" form:"student_id"`
	Course string `json:"course" form:"course" binding:"required"`
	CourseId string `json:"course_id" form:"course_id"`
	TeacherId string `json:"teacher_id" form:"teacher_id"`

}