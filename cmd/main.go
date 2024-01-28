package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jonascavalcanti/square-backend/internal/app"
	"github.com/jonascavalcanti/square-backend/internal/database"
	"github.com/jonascavalcanti/square-backend/internal/delivery/http"
)

func main() {
	r := gin.Default()

	database.InitDB()

	userRepo := database.NewUserRepository(database.DB)
	userService := app.NewUserApp(userRepo)
	userHandler := http.NewUserHandler(userService)

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", userHandler.CreateUser)
			users.GET("/", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	r.Run(":8080")
}
