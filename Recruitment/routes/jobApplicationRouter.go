package routes

import (
	//"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/middleware"
	controller "github.com/Aditya-sairam/golang-jwt-project/Recruitment/controllers"
	"github.com/gin-gonic/gin"
)

func JobAppRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("jobs/job_application", controller.JobApplicationProcess())

}
