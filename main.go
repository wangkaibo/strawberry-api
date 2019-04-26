package main

import (
	"github.com/joho/godotenv"
	"log"
	"strawberry-wallpaper/bootstrap"
	"strawberry-wallpaper/routers"
	"github.com/gin-gonic/gin"
	"os"
)
func init() {
	// 加载环境变量
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	gin.SetMode(os.Getenv("RUN_MODE"))
	app := bootstrap.NewApp("strawberry")
	app.Configure(routers.SetupRouter)
	// Listen and Server in 0.0.0.0:8080
	app.Engine.Run(":8080")
}