package controller

import (
	"CSA_Final_Work/dao"
	"CSA_Final_Work/jwt"
	"CSA_Final_Work/service"
	"CSA_Final_Work/verify"
	"github.com/gin-gonic/gin"
)


func TeacherLogin(c *gin.Context){
	name:=c.Query("name")
	password:=c.Query("password")
	teacher,err:=service.TeacherLogin(name,password)
	if err!=nil{
		c.JSON(500,err.Error())
	}
	token:=jwt.GenToken(teacher.Name,"teacher",teacher.TeacherId)
	c.JSON(200,gin.H{
		"code":200,
		"msg":"登录成功",
		"role":"teacher",
		"token":token,
	})
}


func TeacherChangePassword1(c *gin.Context)  {
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

func TeacherChangePassword2(c *gin.Context) {
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
		service.TeacherChangePasswordByEmail(email,newPassword)
	}
}

func TeacherCheckStudent(c *gin.Context){
	teacherId,_:=c.Get("id")
	students,err:=service.TeacherCheckStudent(teacherId.(string))
	if err != nil {
		c.JSON(500,gin.H{
			"code":500,
			"msg":err.Error(),
		})
	}else{
		c.JSON(200,gin.H{
			"code":200,
			"msg":students,
		})
	}
}
