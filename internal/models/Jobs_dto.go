package models

import "time"

type AddJobDto struct {
	UserId int64 `json:"userid" validate:"required,gte=0"`
	TaskId int64 `json:"taskid" validate:"required,gte=0"`
}

type TrackTimeDto struct {
	UserId int64
	TaskId int64
}
type JobIntervalDto struct {
	DateStart time.Time `validate:"omitempty,validateDate"`
	DateEnd   time.Time `validate:"omitempty,validateDate"`
}
