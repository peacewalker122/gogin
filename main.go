package main

import (
	"gogin/datareq"
	"gogin/handle"
	"gogin/signup"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=bob password=admin dbname=goapi port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err.Error())
	}

	log.Println("DATABASE connection establish")
	db.AutoMigrate(&datareq.DBdata{})
	db.AutoMigrate(&signup.Signup{})

	t := datareq.Newrepo(db)
	ser := datareq.Newservice(t)
	newhandle := handle.Newhandle(ser)
	
	s := signup.Newrepo(db)
	sign := signup.Newservice(s)
	signup := handle.Newsignup(sign)

	r := gin.Default()
	//r.GET("/", newhandle.Getall)
	r.GET("/:id", newhandle.Getbyid)
	r.POST("/", newhandle.Create)
	r.PUT("/:id", newhandle.Update)
	r.DELETE("/:id", newhandle.Delete)
	r.POST("/signup", signup.Signup)
	r.POST("/login",signup.Login)
	r.Run()
}
