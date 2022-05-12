package main

import (
	"CSA_Final_Work/dao"
	"CSA_Final_Work/router"
)

func main() {
	err:=dao.SqlInit()
	if err!=nil{
		panic(err)
	}
	defer dao.SqlClose()
	err=dao.RedisInit()
	if err != nil {
		panic(err)
	}
	defer dao.RedisClose()
	r:=router.SetRouter()
	r.Run(":8080")
}
