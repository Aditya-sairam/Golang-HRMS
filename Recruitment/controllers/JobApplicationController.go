package controllers

import (
	"context"
	"net/http"
	"time"

	"log"

	"github.com/Aditya-sairam/golang-jwt-project/Jwt-Authentication/databases"
	"github.com/Aditya-sairam/golang-jwt-project/Recruitment/models"
	services "github.com/Aditya-sairam/golang-jwt-project/Recruitment/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var JobApplicationCollection *mongo.Collection = databases.OpenCollection(databases.Client, "jobApplications")

func JobApplicationProcess() gin.HandlerFunc {
	return func(c *gin.Context) {
		jobId := c.PostForm("job_id")
		applicantName := c.PostForm("applicant_name")
		phoneNumber := c.PostForm("phone_number")
		email := c.PostForm("email")

		objectId, err := primitive.ObjectIDFromHex(jobId)
		if err != nil {
			log.Printf("Invalid Job Id: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Job Id"})
			return
		}

		resumeFile, err := c.FormFile("resume")
		if err != nil {
			log.Printf("Failed to get resume file: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload resume"})
			return
		}

		coverLetterFile, err := c.FormFile("cover_letter")
		if err != nil {
			log.Printf("Failed to get cover letter file: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload cover letter"})
			return
		}

		resumeURL, err := services.UploadToS3(resumeFile, "resumes/")
		if err != nil {
			log.Printf("Failed to upload resume to S3: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload resume to S3"})
			return
		}

		coverLetterURL, err := services.UploadToS3(coverLetterFile, "cover_letters/")
		if err != nil {
			log.Printf("Failed to upload cover letter to S3: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload cover letter to S3"})
			return
		}

		jobApplication := models.JobApplication{
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
			log.Printf("Error inserting job application into MongoDB: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error applying to this job!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Job application submitted successfully!"})
	}
}
