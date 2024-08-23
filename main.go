package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/routes"
	leaveRoutes "github.com/Aditya-sairam/golang-jwt-project/Leave-Application/routes"
	recruitmentRoutes "github.com/Aditya-sairam/golang-jwt-project/Recruitment/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // React app origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "token"},
		AllowCredentials: true,
	}))

	// Setup routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	leaveRoutes.LeaveAppRoutes(router)
	recruitmentRoutes.JobRoutes(router)
	recruitmentRoutes.JobAppRoutes(router)

	// Print routes for debugging
	for _, route := range router.Routes() {
		fmt.Println(route.Method, route.Path)
	}

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
