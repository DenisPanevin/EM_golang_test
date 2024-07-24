package models

import "time"

type Job struct {
	Id      int64
	UserId  int64
	TaskId  int64
	Started time.Time
	Stoped  time.Time
}
