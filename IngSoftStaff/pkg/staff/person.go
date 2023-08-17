package staff

import "time"


const (
	Student  = 1
	Bachelor = 2
	Graduate = 3
)

const (
	Active     = 1
	Resigned   = 2
	Onboarding = 3
)

type Person struct {
	Id            string
	FirstName     string
	Surname       string
	WorkEmail     string
	PersonalEmail string
	CellPhone     string
	Graduation    string
	Degree        string
	ActiveSince   time.Time
	Status        int
}

type Course struct {
	Id string
	Name     string
}

type Session struct {
	Id   string
	Name        string
	StartDate   time.Time
	Instructors []Person
}

func NewPerson(id, firstName, surname, workEmail, personalEmail, cellPhone, graduation, degree string, activeSince time.Time, status int) Person {
	return Person{
		Id:            id,
		FirstName:     firstName,
		Surname:       surname,
		WorkEmail:     workEmail,
		PersonalEmail: personalEmail,
		CellPhone:     cellPhone,
		Graduation:    graduation,
		Degree:        degree,
		ActiveSince:   activeSince,
		Status:        status,
	}
}
