package model


//用于记录老师和课程的关系
type Tc struct {
	CourseId string `json:"course_id" form:"course_id"`
	CourseName string `json:"course_name" form:"course_name"`
	TeacherId string `json:"teacher_id" form:"teacher_id"`
}
