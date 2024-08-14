package todo

import (
	"strings"
	"time"
)

type TODO struct {
	Description string
	Date        time.Time
}

func NewTODO(Desc string, Date string, Time string) TODO {

	datetime, _ := time.Parse(time.DateTime, strings.Join([]string{Date, Time}, " "))

	return TODO{
		Description: Desc,
		Date:        datetime,
	}

}
