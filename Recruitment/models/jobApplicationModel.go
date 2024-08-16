package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// JobApplication represents a user's application to a job
type JobApplication struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	JobID          string             `json:"job_id"` // Reference to JobPosting
	ApplicantName  string             `bson:"applicant_name" json:"applicant_name"`
	Email          string             `bson:"email" json:"email"`
	PhoneNumber    string             `bson:"phone_number" json:"phone_number"`
	ResumeURL      string             `bson:"resume_url" json:"resume_url"`             // Link to resume in S3
	CoverLetterURL string             `bson:"cover_letter_url" json:"cover_letter_url"` // Link to cover letter in S3
	AppliedDate    time.Time          `bson:"applied_date" json:"applied_date"`
	Status         string             `bson:"status" json:"status"` // Status of the application (e.g., Submitted, Reviewed)
}
