package staff

import "time"

/*
type PersonDTO struct {
    ID       string `json:"id" csv:"id"`
    Name     string `json:"name" csv:"name"`
    Age      int    `json:"age" csv:"age"`
    Location string `json:"location" csv:"location"`
}
*/



const (
    Student = 1
    Bachelor = 2
    Graduate = 3
)


type PersonDTO struct {
	IDPersona string
	FirstName string
	Surname   string
	WorkEmail string
    PersonalEmail string
    CellPhone string
    Graduation string
    Degree string
}

type CourseDTO struct {
    IDCourse string
    Name     string
}

type SessionDTO struct {
    IDSession string
    Name      string
    StartDate time.Time
    Instructors []PersonDTO
}
