package todo

import (
	"time"
)

type TODO struct {
	Description string
	Day         time.Time
	Time        time.Time
}

func NewTODO(Desc string, Date string, Time string) TODO {

	day, _ := time.Parse(time.DateOnly, Date)
	time, _ := time.Parse(time.TimeOnly, Time)

	return TODO{
		Description: Desc,
		Day:         day,
		Time:        time,
	}

}
