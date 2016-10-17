package events

import (
	"time"

	"github.com/mkappus/checkin/src/users"
)

type (
	Checkin struct {
		ID       int
		User     *users.Student
		Loc      string
		Start    time.Time
		End      time.Time
		IsActive bool
	}
)
