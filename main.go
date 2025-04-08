package main

import (
	"log"
	"net/http"
	"warehouse-backend/config"
	"warehouse-backend/handlers"
	"warehouse-backend/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}
	db := config.InitDB()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	authHandler := handlers.NewAuthHandler(db)
	barangHandler := handlers.NewBarangHandler(db)
	r.POST("/login", authHandler.Login)
	protected := r.Group("/")
	protected.Use(middleware.JWTAuth())
	{
		protected.POST("/barang/create", barangHandler.CreateBarang)
		protected.GET("/barang/get", barangHandler.GetAllBarang)
		protected.PUT("/barang/update/:id", barangHandler.UpdateBarang)
		protected.DELETE("/barang/delete/:id", barangHandler.DeleteBarang)
	}
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}

}
