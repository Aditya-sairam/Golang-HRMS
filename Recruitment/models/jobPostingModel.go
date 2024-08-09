package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobPosting struct {
	JobId                primitive.ObjectID `bson:"_id"`
	JobTitle             string             `json:"job_name" validate:"required"`
	Description          string             `json:"description" validate:"required,min=2,max=300"`
	PreferredSkills      string             `json:"preferred_skills"`
	Status               string             `json:"status" validate:"required,oneof=Hired Active Inactive InProgress"`
	PostedDate           time.Time          `json:"posted_date"`
	LastUpdated          time.Time          `json:"last_updated"`
	NumberOfApplications int                `json:"number_of_applications"`
	Type                 string             `json:"type" validate:"required,oneof=Internal External"`
	Department           string             `json:"department"`
	Location             string             `json:"location"`
	Job_Id               string             `json:"job_id"`
}
