package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// public routes
	r.GET("/articles", handlers.GetArticles)
	r.GET("/articles/:id", handlers.GetArtcile)

	// login route
	r.POST("/login", handlers.Login)

	// admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthRequired())
	{
		admin.POST("/articles", handlers.CreateArticle)
		admin.PUT("/articles/:id", handlers.UpdateArticle)
		admin.DELETE("/articles/:id", handlers.DeleteArticle)
	}

	// launch server
	r.Run(":8080")
}
