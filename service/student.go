package service

import (
	"CSA_Final_Work/dao"
	"CSA_Final_Work/model"
	"errors"
)

func StudentLogin(name string,password string)(*model.Student,error)  {
	student:=new(model.Student)
	if err:=dao.DB.Where("name=? and password=?",name,password).First(student).Error;err!=nil{
		return nil,errors.New("账号或密码错误")
	}
	return student,nil
}

//通过邮箱改密码
func StudentChangePasswordByEmail(email string,newPassword string){
	student:=new(model.Student)
	dao.DB.Model(student).Where("email = ?",email).Update("password",newPassword)
	dao.RD.Del(email)
}

func StudentCheckCourse(studentId string)(*[]model.Record,error){
	records:=new([]model.Record)
	err:=dao.DB.Table("records").Where("student_id = ?",studentId).Find(records).Error
	if err != nil {
		return nil,errors.New("未找到记录")
	}
	return records,nil
}

func StudentChooseCourse(name string,studentId string,courseId string)error  {
	course:=new(model.Course)
	if err:=dao.DB.Table("courses").Where("course_id =?",courseId).First(course).Error;err != nil {
		return errors.New("课程不存在")
	}
	choseCourse :=new(model.Course)
	records:=new([]model.Record)
	if err:=dao.DB.Table("records").Where("student_id =?",studentId).Find(records).Error;err==nil {
		for _, record := range *records {
			dao.DB.Table("courses").Where("course_id=?", record.CourseId).First(choseCourse)
			if choseCourse.Weekday == course.Weekday && choseCourse.Class==course.Class{
				return errors.New("课程冲突，选课失败")
			}
		}
	}
	if course.Max<=course.Now{
		return errors.New("课程人数已满")
	}
	record:=&model.Record{
		Student:   name,
		StudentId: studentId,
		Course:    course.Name,
		CourseId:  courseId,
		TeacherId: course.TeacherId,
	}

	tx:=dao.DB.Begin()
	if err:=dao.DB.Save(record).Error;err!=nil{
		tx.Rollback()
		return errors.New("选课失败，请重试")
	}
	if err:=dao.DB.Model(course).Update("now",course.Now+1).Error;err!=nil{
		tx.Rollback()
		return errors.New("选课失败，请重试")
	}
	tx.Commit()
	return nil
}
