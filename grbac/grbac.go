package grbac

import (
	"github.com/gin-gonic/gin"
	"github.com/storyicon/grbac"
	"time"
)


//grbac
func GrbacAuthorization() gin.HandlerFunc {
	rbac,err:=grbac.New(grbac.WithJSON("grbac.json",time.Minute*5))
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		roles:=make([]string,5)
		role,_:=c.Get("role")
		roles=append(roles,role.(string))
		state,err:=rbac.IsRequestGranted(c.Request,roles)
		if err != nil {
			c.JSON(500,gin.H{
				"code":500,
				"msg":"权限未知，访问失败",
			})
			c.Abort()
			return
		}
		if !state.IsGranted(){
			c.JSON(500,gin.H{
				"code":500,
				"msg":"权限不足，访问失败",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}


