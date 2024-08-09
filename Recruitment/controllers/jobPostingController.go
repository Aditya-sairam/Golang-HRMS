package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/databases"
	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/helpers"
	"github.com/Aditya-sairam/golang-jwt-project/Recruitment/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var jobPostingCollection *mongo.Collection = databases.OpenCollection(databases.Client, "user")

func CreateJobPosting() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jobPosting models.JobPosting
		err := helpers.CheckUserType(c, "ADMIN")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "This page can only be accessed by admins!"})
			return
		}

		err = c.BindJSON(&jobPosting)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

		validateError := validate.Struct(jobPosting)

		if validateError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": validateError})
			return
		}

		jobPosting.JobId = primitive.NewObjectID()
		jobPosting.PostedDate, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		jobPosting.LastUpdated, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		resultInserstionNumber, insertionErr := jobPostingCollection.InsertOne(ctx, jobPosting)
		if insertionErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": insertionErr})
		}
		defer cancel()
		c.JSON(http.StatusOK, resultInserstionNumber)

	}
}
