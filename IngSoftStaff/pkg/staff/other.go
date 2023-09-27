package staff
import ("time")

type Course struct {
	Id   string
	Name string
}

type Session struct {
	Id          string
	Name        string
	StartDate   time.Time
	Instructors []PersonData
}
