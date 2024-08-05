package models

type LeaveType struct {
	LeaveTypeID   int    `bson:"_id"`
	LeaveTypeName string `json:"leave_type" validate:"required,eq=Sick|eq=Casual|eq=Paid|eq=Maternity"`
	Leave_id      string `json:"leave_id"`
}
