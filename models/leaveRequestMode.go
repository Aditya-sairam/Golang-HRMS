package models

import "time"

type LeaveRequest struct {
	LeaveRequestID int       `bson:"_id"`
	UserId         string    `json:"user_id"`
	LeaveTypeID    string    `json:"leave_id"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"End_Date"`
	Reason         string    `json:"reason"`
	Status         string    `json:"status" validate:"required,eq=Pending|eq=Approved|eq=Denied"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
