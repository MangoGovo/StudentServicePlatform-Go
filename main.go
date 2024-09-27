package main

import (
	"StuService-Go/internal/middleware"
	database "StuService-Go/internal/pkg/databse"
	"StuService-Go/internal/router"
	"StuService-Go/internal/service"
	"StuService-Go/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	utils.InitLogger()
	db := database.Init()
	service.Init(db)
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.ErrHandler())
	r.NoMethod(middleware.HandleNotFond)
	r.NoRoute(middleware.HandleNotFond)
	r.Static("/static", "./static")
	router.Init(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
