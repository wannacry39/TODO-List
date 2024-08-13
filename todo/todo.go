package todo

type TODO struct {
	Description string
	Date        string
	Time        string
}

func NewTODO(Desc string, Date string, Time string) TODO {
	return TODO{
		Description: Desc,
		Date:        Date,
		Time:        Time,
	}

}
