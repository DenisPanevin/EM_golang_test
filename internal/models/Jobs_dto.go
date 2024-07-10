package models

type AddJobDto struct {
	UserId    int64 `json:"userid" validate:"required,gte=0"`
	TaskId    int64 `json:"taskid" validate:"required,gte=0"`
	StartStop bool  `json:"startstop" validate:"required"`
}

type TrackTimeDto struct {
	UserId int64
	TaskId int64
}
