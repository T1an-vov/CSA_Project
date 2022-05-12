package controller

import (
	"CSA_Final_Work/dao"
	"CSA_Final_Work/jwt"
	"CSA_Final_Work/service"
	"CSA_Final_Work/verify"
	"github.com/gin-gonic/gin"
)

func StudentLogin(c *gin.Context){
	name:=c.Query("name")
	password:=c.Query("password")
	student,err:=service.StudentLogin(name,password);
	if err!=nil{
		c.JSON(500,err.Error())
	}
	token:=jwt.GenToken(student.Name,"student",student.StudentId)
	c.JSON(200,gin.H{
		"code":200,
		"msg":"登录成功",
		"role":"student",
		"token":token,
	})
}


func StudentChangePassword1(c *gin.Context)  {
	email:=c.Query("email")
	str:=verify.GetRandomString()
	if err:=verify.SendEmail(email,str);err!=nil{
		c.JSON(500,gin.H{
			"code":500,
			"msg":err.Error(),
		})
	}else{
		c.JSON(200,gin.H{
			"code":200,
			"msg":"发送验证码成功，有效时间为10分钟",
		})
	}
}

func StudentChangePassword2(c *gin.Context) {
	email:=c.PostForm("email")
	newPassword:=c.PostForm("newpassword")
	verifyCode:=c.PostForm("verifycode")
	code,err:=dao.RD.Get(email).Result()
	if err != nil {
		c.JSON(500,gin.H{
			"code":500,
			"msg":"验证码已失效",
		})
	}else if code!=verifyCode {
		c.JSON(500,gin.H{
			"code":500,
			"msg":"验证码错误",
		})
	}else{
		service.StudentChangePasswordByEmail(email,newPassword)
	}
}


func StudentCheckCourse(c *gin.Context) {
	studentId,_:=c.Get("id")
	records,err:=service.StudentCheckCourse(studentId.(string))
	if err != nil {
		c.JSON(200,gin.H{
			"code":200,
			"msg":err.Error(),
		})
	}else{
		c.JSON(200,gin.H{
			"code":200,
			"msg":records,
		})
	}
}


func StudentChooseCourse(c *gin.Context)  {
	studentId,_:=c.Get("id")
	name,_:=c.Get("name")
	courseId:=c.PostForm("course_id")
	err:=service.StudentChooseCourse(name.(string),studentId.(string),courseId)
	if err != nil {
		c.JSON(500,gin.H{
			"code":500,
			"msg":err.Error(),
		})
	}else{
		c.JSON(200,gin.H{
			"code":200,
			"msg":name.(string)+"选择"+courseId+"成功",
		})
	}
}

