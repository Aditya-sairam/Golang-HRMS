package routes

import (
	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/middleware"
	controller "github.com/Aditya-sairam/golang-jwt-project/Recruitment/controllers"
	"github.com/gin-gonic/gin"
)

func JobRoutes(incomingRoutes *gin.Engine) {

	authenticatedRoutes := incomingRoutes.Group("/recruitment")
	authenticatedRoutes.Use(middleware.Authenticate())

	authenticatedRoutes.GET("/:job_id/applicants_list", controller.ListApplicants())
	// Routes that require authentication

	authenticatedRoutes.POST("job_posting", controller.CreateJobPosting())
	authenticatedRoutes.GET("job_list", controller.ListJobs())
	authenticatedRoutes.GET("/:job_id", controller.GetJob())

}
