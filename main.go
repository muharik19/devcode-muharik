package main

import (
	"fmt"
	"os"

	controller "github.com/devcode-muharik/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	port := os.Getenv("APP_PORT")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// ACTIVITY
	r.POST("activity-groups", controller.CreatedActivity)
	r.PATCH("activity-groups/:id", controller.UpdateActivity)
	r.DELETE("activity-groups/:id", controller.DeleteActivity)
	r.GET("activity-groups", controller.ListActivityAll)
	r.GET("activity-groups/:id", controller.ListActivityDetail)
	// TODO
	r.POST("todo-items", controller.CreatedTodo)
	r.PATCH("todo-items/:id", controller.UpdateTodo)
	r.DELETE("todo-items/:id", controller.DeleteTodo)
	r.GET("todo-items", controller.ListTodoAll)
	r.GET("todo-items/:id", controller.ListTodoDetail)
	fmt.Println("Staring server port " + port + "")
	r.Run(":" + port)
}
