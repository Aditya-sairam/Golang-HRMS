package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LeaveRequest struct {
	LeaveRequestID primitive.ObjectID `bson:"_id"`
	UserId         string             `json:"user_id"`
	LeaveTypeID    string             `json:"leave_id"`
	StartDate      time.Time          `json:"start_date"`
	LeaveTypeName  string             `json:"leave_type" validate:"required,eq=Sick|eq=Casual|eq=Paid|eq=Maternity"`
	EndDate        time.Time          `json:"End_Date"`
	Reason         string             `json:"reason"`
	Status         string             `json:"status" validate:"required,eq=Pending|eq=Approved|eq=Denied"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}
