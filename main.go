package main

import (
	"StudentServicePlatform_Go/internal/middleware"
	database "StudentServicePlatform_Go/internal/pkg/databse"
	"StudentServicePlatform_Go/internal/router"
	"StudentServicePlatform_Go/internal/service"
	"StudentServicePlatform_Go/pkg/utils"
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
	r.NoMethod(middleware.HandleNotFond)
	r.NoRoute(middleware.HandleNotFond)
	router.Init(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
