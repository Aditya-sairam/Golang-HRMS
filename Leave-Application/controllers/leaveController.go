package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/databases"
	"github.com/Aditya-sairam/golang-jwt-project/Leave-Application/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"golang.org/x/crypto/bcrypt"
)

var leaveValidate = validator.New()

var leaveCollection *mongo.Collection = databases.OpenCollection(databases.Client, "user")

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
		defer cancel()
		leaveRequest.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		leaveRequest.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		uid, _ := c.Get("uid")
		uidStr, _ := uid.(string)
		leaveRequest.UserId = uidStr
		//user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		leaveRequest.LeaveRequestID = primitive.NewObjectID()
		leaveRequest.LeaveTypeID = leaveRequest.LeaveRequestID.Hex()
		leaveRequest.Status = "Pending"
		resultInserstionNumber, insertErr := leaveCollection.InsertOne(ctx, leaveRequest)
		if insertErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": insertErr})
			c.Abort()
		}
		c.JSON(http.StatusOK, resultInserstionNumber)

	}
}
