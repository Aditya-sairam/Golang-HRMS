package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/databases"
	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/helpers"
	"github.com/Aditya-sairam/golang-jwt-project/Leave-Application/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var leaveValidate = validator.New()
var leaveCollection *mongo.Collection = databases.OpenCollection(databases.Client, "leaveApplication")

func LeaveRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		var leaveRequest models.LeaveRequest
		if err := c.BindJSON(&leaveRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		leaveRequest.Status = "Pending"
		if err := leaveValidate.Struct(leaveRequest); err != nil {
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.StructNamespace()+": "+err.ActualTag()+" failed on "+err.Param())
			}
			c.JSON(http.StatusBadRequest, gin.H{"validation_errors": errors})
			return
		}

		leaveRequest.CreatedAt = time.Now()
		leaveRequest.UpdatedAt = time.Now()

		uid, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found"})
			return
		}

		uidStr, ok := uid.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id is of invalid type"})
			return
		}

		leaveRequest.UserId = uidStr
		leaveRequest.LeaveRequestID = primitive.NewObjectID()
		leaveRequest.LeaveTypeID = leaveRequest.LeaveRequestID.Hex()

		resultInsertionNumber, insertErr := leaveCollection.InsertOne(ctx, leaveRequest)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
			return
		}

		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func LeaveList() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := helpers.CheckUserType(c, "ADMIN")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "This page can only be accessed by admins!"})
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 60*time.Second)
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}
		//startIndex := (page-1)*recordPerPage
		startIndex, err := strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{
			{"_id", bson.D{{"_id", "null"}}},
			{"total_count", bson.D{{"$sum", 1}}},
			{"data", bson.D{{"$push", "$$ROOT"}}}}}}

		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			}},
		}
		result, err := leaveCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		var allUsers []bson.M
		if err = result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allUsers[0])
	}
}
