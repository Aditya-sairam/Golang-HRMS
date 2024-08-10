package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/databases"
	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/helpers"
	"github.com/Aditya-sairam/golang-jwt-project/Recruitment/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Please add comments to all the file (Good coding practice)
var validate = validator.New()
var jobPostingCollection *mongo.Collection = databases.OpenCollection(databases.Client, "jobs")

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
		jobPosting.Job_Id = jobPosting.JobId.Hex()
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

func ListJobs() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		result, err := jobPostingCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while listing user items"})
			return
		}
		var allPosting []bson.M
		if err = result.All(ctx, &allPosting); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allPosting[0])
	}
}

func GetJob() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("job_id")

		err := helpers.MatchUserTypeToUid(c, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.JobPosting
		err = jobPostingCollection.FindOne(ctx, bson.M{"job_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
