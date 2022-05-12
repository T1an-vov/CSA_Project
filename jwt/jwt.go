package jwt

import (
	"CSA_Final_Work/dao"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type MyClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Id string `json:"id"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Minute*10
var MySecret = []byte("981774028")

//生成JWT
func GenToken(name string,role string,id string)string{
	c:=MyClaims{
		Id: id,
		Name: name,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 	time.Now().Add(TokenExpireDuration).Unix(),	//token的过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)	// 使用签名方法创建签名对象
	tokenstring,_ := token.SignedString(MySecret)
	dao.RD.Do("Set",tokenstring,role)
	dao.RD.Do("expire",tokenstring,c.ExpiresAt-time.Now().Unix())
	return tokenstring					//使用指定的secret签名并获得完整的编码后的字符串token
}

//解析token
func parseToken(tokenString string) *MyClaims {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		panic("token is valid")
	}
	return claims
}

//JWT认证
func JWTAuthMiddleware()func(c *gin.Context){
	return func(c *gin.Context){
		//客户端携带token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader ==""{
			c.JSON(500,gin.H{
				"code":500,
				"msg":"token为空，请求失败，请先登录",
			})
			c.Abort()
			return
		}
		//按空格进行分割判断请求头是否正确
		parts := strings.SplitN(authHeader," ",2)
		if !(len(parts)==2&&parts[0]=="Bearer"){
			c.JSON(500,gin.H{
				"code":500,
				"msg":"token 错误",
			})
			c.Abort()
			return
		}

		//从redis数据库中判断token是否有效
		_,err := dao.RD.Get(parts[1]).Result()
		if err != nil {
			c.JSON(500,gin.H{
				"code":500,
				"msg":"token 过期",
			})
			c.Abort()
			return
		}
		claim:=parseToken(parts[1])
		c.Set("name",claim.Name)
		c.Set("role",claim.Role)
		c.Set("id",claim.Id)
		c.Next()
	}
}