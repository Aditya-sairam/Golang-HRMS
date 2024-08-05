package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Aditya-sairam/golang-jwt-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var leaveValidate = validator.New()

func LeaveRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 60*time.Second)
		var leaveRequest models.LeaveRequest
		err := c.BindJSON(&leaveRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		validateError := leaveValidate.Struct(leaveRequest)
		if validateError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Please check the details you have entered!"})
		}

	}
}
