package staff

import (

	"time"
	//"IngSoftStaff/pkg/util"
)


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

type PersonData struct {
	ID            string
	FirstName     string
	Surname       string
	WorkEmail     string
	PersonalEmail string
	CellPhone     string
	Graduation    string
	Degree        string
	LastPromotion string
	ActiveSince   string
}



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

func NewPerson(id, firstName, surname, workEmail, personalEmail, cellPhone, graduation, degree string, lastPrometed time.Time, activeSince time.Time, status int) PersonData {
	return PersonData{
		ID:            id,
		FirstName:     firstName,
		Surname:       surname,
		WorkEmail:     workEmail,
		PersonalEmail: personalEmail,
		CellPhone:     cellPhone,
		Graduation:    graduation,
		Degree:        degree,
		LastPromotion: lastPrometed.String(),
		ActiveSince:   activeSince.String(),
	}
}



