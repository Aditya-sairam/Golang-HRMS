package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Aditya-sairam/golang-jwt-project/Recruitment/models"
	services "github.com/Aditya-sairam/golang-jwt-project/Recruitment/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var jobApplication models.JobApplication
var JobApplicationCollection *mongo.Collection

func JobApplicationProcess() gin.HandlerFunc {
	return func(c *gin.Context) {

		jobId := c.PostForm("job_id")
		applicantName := c.PostForm("applicant_name")
		phoneNumber := c.PostForm("phone_number")
		email := c.PostForm("email")
		objectId, err := primitive.ObjectIDFromHex(jobId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Job Id"})
		}

		resumeFile, err := c.FormFile("resume")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload resume"})
			return
		}

		coverLetterFile, err := c.FormFile("cover_letter")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload cover letter"})
			return
		}
		resumeURL, err := services.UploadToS3(resumeFile, "resumes/")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload resume to S3"})
			return
		}

		coverLetterURL, err := services.UploadToS3(coverLetterFile, "cover_letters/")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload cover letter to S3"})
			return
		}
		jobApplication = models.JobApplication{
			ID:             primitive.NewObjectID(),
			JobID:          objectId,
			ApplicantName:  applicantName,
			Email:          email,
			PhoneNumber:    phoneNumber,
			ResumeURL:      resumeURL,
			CoverLetterURL: coverLetterURL,
			AppliedDate:    time.Now(),
			Status:         "Submitted",
		}
		_, err = JobApplicationCollection.InsertOne(context.TODO(), jobApplication)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error applying to this job!"})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Job APplication submitted successfully!"})
	}
}
