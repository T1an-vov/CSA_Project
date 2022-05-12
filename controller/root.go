package controller

import (
	"CSA_Final_Work/jwt"
	"CSA_Final_Work/model"
	"CSA_Final_Work/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = *validator.New()

func RootLogin(c *gin.Context){
	name:=c.Query("name")
	password:=c.Query("password")
	_,err:=service.RootLogin(name,password)
	if err!=nil{
		c.JSON(500,err.Error())
	}else{
	token:=jwt.GenToken("","root","")
	c.JSON(200,gin.H{
		"code":200,
		"msg":"登录成功",
		"role":"root",
		"token":token,
	})}
}


func AddTeacher(c *gin.Context){
	teacher:=new(model.Teacher)
	err:=c.ShouldBind(teacher)
	if err!=nil{
		c.JSON(500,gin.H{
			"code":500,
			"msg":"参数错误，创建老师失败",
		})
	}else {
		if err := validate.Struct(teacher); err != nil {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "email错误",
			})
		} else {
			err = service.AddTeacher(teacher)
			if err != nil {
				c.JSON(500, gin.H{
					"code": 500,
					"err":  err.Error(),
				})
			} else {
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "创建老师成功",
				})
			}
		}
	}
}


func AddStudent(c *gin.Context){
	student:=new(model.Student)
	err:=c.ShouldBind(student)
	if err!=nil{
		c.JSON(500,gin.H{
			"code":500,
			"msg":"参数错误，创建学生失败",
		})
	}else {
		if err := validate.Struct(student); err != nil {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  "email错误",
			})
		} else {
			err = service.AddStudent(student)
			if err != nil {
				c.JSON(500, gin.H{
					"code": 500,
					"err":  err.Error(),
				})
			} else {
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "创建学生成功",
				})
			}
		}
	}
}


func DeleteTeacher(c *gin.Context){
	teacherId :=c.PostForm("teacher_id")
	if err:=service.DeleteTeacher(teacherId);err!=nil{
		c.JSON(500,gin.H{
			"code":500,
			"msg":err.Error(),
		})
	}else{
		c.JSON(200,gin.H{
			"code":200,
			"msg":"删除老师成功",
		})
	}
}


func DeleteStudent(c *gin.Context){
	studentId:=c.PostForm("student_id")
	if err:=service.DeleteStudent(studentId);err!=nil{
		c.JSON(500,gin.H{
			"code":500,
			"msg":err.Error(),
		})
	}else{
		c.JSON(200,gin.H{
			"code":200,
			"msg":"删除学生成功",
		})
	}
}


func AddCourse(c *gin.Context){
	course:=new(model.Course)
	err:=c.ShouldBind(course)
	if err!=nil{
		c.JSON(500,gin.H{
			"code":500,
			"msg":"参数错误，创建课程失败",
		})
	}else {
		err = service.AddCourse(course)
		if err != nil {
			c.JSON(500, gin.H{
				"code": 500,
				"err":  err.Error(),
			})
		}else{
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "创建课程成功",
			})
		}
	}
}


func UpdateCourse(c *gin.Context){
	CourseId:=c.PostForm("course_id")
	newCourse:=new(model.Course)
	err:=c.ShouldBind(newCourse)
	if err!=nil{
		c.JSON(500,gin.H{
			"code":500,
			"msg":"参数错误，更新课程失败",
		})
	}else {
		err = service.UpdateCourse(CourseId,newCourse)
		if err != nil {
			c.JSON(500, gin.H{
				"code": 500,
				"err":  err.Error(),
			})
		}else{
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "更新课程成功",
			})
		}
	}
}
