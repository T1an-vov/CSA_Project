package service

import (
	"CSA_Final_Work/dao"
	"CSA_Final_Work/model"
	"errors"
)

//登录
func TeacherLogin(name string, password string) (*model.Teacher,error) {
	teacher:=new(model.Teacher)
	if err:=dao.DB.Where("name=? and password=?",name,password).First(teacher).Error;err!=nil{
		return nil,errors.New("账号或密码错误")
	}
	return teacher,nil
}

//通过邮箱改密码
func TeacherChangePasswordByEmail(email string,newPassword string)  {
	teacher:=new(model.Teacher)
	dao.DB.Model(teacher).Where("email = ?",email).Update("password",newPassword)
	dao.RD.Del(email)
}


//老师查看选择自己课程的学生
func TeacherCheckStudent(teacherId string) (student []string,err error) {
	records:=new([]model.Record)
	err=dao.DB.Table("records").Where("teacher_id = ?",teacherId).Find(records).Error
	if err != nil {
		return nil,errors.New("未找到记录")
	}
	for _, record := range *records {
		student=append(student,record.Student)
	}
	return student,nil
}