package todo

type TODO struct {
	Description string
	Day         string
	Time        string
}

func NewTODO(Desc string, Date string, Time string) TODO {

	return TODO{
		Description: Desc,
		Day:         Date,
		Time:        Time,
	}

}
