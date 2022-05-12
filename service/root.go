package service

import (
	"CSA_Final_Work/dao"
	"CSA_Final_Work/model"
	"errors"
	"strconv"
)

func RootLogin(name,password string)(*model.Root, error) {
	root:=new(model.Root)
	if err:=dao.DB.Where("name=? and password=?",name,password).First(root).Error;err!=nil{
		return nil,errors.New("账号或密码错误")
	}
	return root,nil
}
func AddStudent(student *model.Student) error {
	if err:=dao.DB.Save(student).Error;err!=nil{
		return errors.New("邮箱已被使用，添加学生失败")
	}
	dao.DB.Model(student).Update("student_id",strconv.Itoa(int(student.Id)))
	return nil
}
func AddTeacher(teacher *model.Teacher)error{
	if err:=dao.DB.Save(teacher).Error;err!=nil{
		return errors.New("邮箱已被使用，添加老师失败")
	}
	dao.DB.Model(teacher).Update("teacher_id",strconv.Itoa(int(teacher.Id)))
	return nil
}

func DeleteTeacher(teacherId string) error {
	teacher:=new(model.Teacher)
	if err:=dao.DB.Table("teachers").Where("teacher_id=?",teacherId).First(teacher).Error;err!=nil{
		return errors.New("未找到对应老师，删除老师失败")
	}
	dao.DB.Delete(teacher)
	return nil
}

func DeleteStudent(studentId string)error{
	student:=new(model.Student)
	if err:=dao.DB.Table("students").Where("student_id=?",studentId).First(student).Error;err!=nil{
		return errors.New("未找到对应学生，删除学生失败")
	}
	dao.DB.Delete(student)
	return nil
}

func AddCourse(course *model.Course)error{
	teacher:=new(model.Teacher)
	if err:=dao.DB.Table("teachers").Where("teacher_id = ?",course.TeacherId).First(teacher).Error;err!=nil{
		return errors.New("该老师不存在，创建课程失败")
	}
	//tx:=dao.DB.Begin()TODO
	dao.DB.Save(course)
	dao.DB.Model(course).Update("course_id",strconv.Itoa(int(course.Id)))
	return nil
}

func UpdateCourse(CourseId string,newCourse *model.Course)error{
	course:=new(model.Course)
	teacher:=new(model.Teacher)
	if err:=dao.DB.Table("courses").Where("course_id = ?",CourseId).First(course).Error;err!=nil{
		return errors.New("未找到对应课程，更新课程失败")
	}
	if err:=dao.DB.Table("teachers").Where("teacher_id = ?",newCourse.TeacherId).First(teacher).Error;err!=nil{
		return errors.New("该老师不存在，更新课程失败")
	}
	dao.DB.Model(course).Update(newCourse)
	return nil
}
