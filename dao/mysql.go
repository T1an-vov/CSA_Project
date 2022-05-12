package dao

import (
	"CSA_Final_Work/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func SqlInit() (err error) {
	db, err := gorm.Open("mysql", "root:root@/csa_final_work?charset=utf8mb4")
	if err != nil {
		return err
	}
	err = db.DB().Ping()
	if err != nil {
		return err
	}
	db.AutoMigrate(&model.Student{})
	db.AutoMigrate(&model.Teacher{})
	db.AutoMigrate(&model.Course{})
	db.AutoMigrate(&model.Record{})
	db.AutoMigrate(&model.Root{})
	db.Exec("alter table students auto_increment = 201110001")//统一格式22111xxxx
	db.Exec("alter table teachers auto_increment = 202220001")//统一格式22222xxxx
	db.Exec("alter table courses  auto_increment = 204440001")//统一格式20444xxxx
	DB = db
	return nil
}
func SqlClose() {
	DB.Close()
}