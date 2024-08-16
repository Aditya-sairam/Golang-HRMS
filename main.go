package main

import (
	"fmt"
	"os"

	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/routes"
	leaveRoutes "github.com/Aditya-sairam/golang-jwt-project/Leave-Application/routes"

	recruitmentRoutes "github.com/Aditya-sairam/golang-jwt-project/Recruitment/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	//routes.authRouter(router)
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	leaveRoutes.LeaveAppRoutes(router)

	recruitmentRoutes.JobRoutes(router)
	recruitmentRoutes.JobAppRoutes(router)

	for _, route := range router.Routes() {
		fmt.Println(route.Method, route.Path)
	}
	//leav
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})
	router.Run(":" + port)

}
