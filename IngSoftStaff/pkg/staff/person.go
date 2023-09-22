package staff

import (
	"fmt"
	"time"
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

type Person struct {
	Id            string `csv:"Persona"`
	FirstName     string `csv:"Nombre"`
	Surname       string `csv:"Apellido"`
	WorkEmail     string `csv:"Email_ORT"`
	PersonalEmail string `csv:"Email_Personal"`
	CellPhone     string `csv:"Celular"`
	Graduation    string `csv:"Graduado"`
	Degree        string `csv:"Graduado"`
}

func (t Person) String() string {
	return fmt.Sprintf("{%v %v }", t.Id, t.FirstName)
}

type Course struct {
	Id   string
	Name string
}

type Session struct {
	Id          string
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
	}
}
