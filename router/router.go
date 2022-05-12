package router

import (
	"CSA_Final_Work/controller"
	"CSA_Final_Work/grbac"
	"CSA_Final_Work/jwt"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r:=gin.Default()
	RootGroup:=r.Group("root")
	{
		RootGroup.GET("login",controller.RootLogin)
		RootGroup.POST("teacher",jwt.JWTAuthMiddleware(),grbac.GrbacAuthorization(),controller.AddTeacher)
		RootGroup.POST("student",jwt.JWTAuthMiddleware(),grbac.GrbacAuthorization(),controller.AddStudent)
		RootGroup.DELETE("teacher",jwt.JWTAuthMiddleware(),grbac.GrbacAuthorization(),controller.DeleteTeacher)
		RootGroup.DELETE("student",jwt.JWTAuthMiddleware(),grbac.GrbacAuthorization(),controller.DeleteStudent)
		RootGroup.POST("course",jwt.JWTAuthMiddleware(),grbac.GrbacAuthorization(),controller.AddCourse)
		RootGroup.PUT("course",jwt.JWTAuthMiddleware(),grbac.GrbacAuthorization(),controller.UpdateCourse)
	}
	TeacherGroup:=r.Group("teacher")
	{
		TeacherGroup.GET("login",controller.TeacherLogin)
		TeacherGroup.GET("password",controller.TeacherChangePassword1)
		TeacherGroup.PUT("password",controller.TeacherChangePassword2)
		TeacherGroup.GET("student",jwt.JWTAuthMiddleware(),grbac.GrbacAuthorization(),controller.TeacherCheckStudent)
	}
	StudentGroup:=r.Group("student")
	{
		StudentGroup.GET("login",controller.StudentLogin)
		StudentGroup.GET("course",jwt.JWTAuthMiddleware(),grbac.GrbacAuthorization(),controller.StudentCheckCourse)
		StudentGroup.POST("course",jwt.JWTAuthMiddleware(),grbac.GrbacAuthorization(),controller.StudentChooseCourse)
		StudentGroup.GET("password",controller.StudentChangePassword1)
		StudentGroup.PUT("password",controller.StudentChangePassword2)
	}
	return r
}
