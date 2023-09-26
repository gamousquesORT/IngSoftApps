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

type DateTime struct {
	time.Time
}

// Convert the internal date as CSV string
func (date *DateTime) MarshalCSV() (string, error) {
	return date.String(), nil
}

// You could also use the standard Stringer interface
func (date DateTime) String() string {
	return date.Time.Format("01022006")
}

// Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("01022006", csv)
	return err
}

type Person struct {
	ID            string `csv:"Persona"`
	FirstName     string `csv:"Nombre"`
	Surname       string `csv:"Apellido"`
	WorkEmail     string `csv:"Email_ORT"`
	PersonalEmail string `csv:"Email_Personal"`
	CellPhone     string `csv:"Celular"`
	Graduation    string `csv:"Graduado"`
	Degree        string `csv:"Grado"`
	LastPromotion string `csv:"FechaPromocion"`
}

func (t Person) String() string {
	return fmt.Sprintf("{%v %v }", t.ID, t.FirstName)
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
		ID:            id,
		FirstName:     firstName,
		Surname:       surname,
		WorkEmail:     workEmail,
		PersonalEmail: personalEmail,
		CellPhone:     cellPhone,
		Graduation:    graduation,
		Degree:        degree,
	}
}
