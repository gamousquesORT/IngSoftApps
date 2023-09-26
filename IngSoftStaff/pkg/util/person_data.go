package util

import (
	"time"
	"IngSoftStaff/pkg/staff"
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

type PersonCSV struct {
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

func NewPersonCSV(id, firstName, surname, workEmail, personalEmail, cellPhone, graduation, degree string, activeSince time.Time, status int) PersonCSV {
	return PersonCSV{
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

type PersonCSVReader interface {
    ReadCSV(filename string) ([]PersonCSV, error)
}

func NewPersonFromCSV(filename string) ([]staff.PersonData, error) {
    csvData, err := ReadCSV(filename)
    if err != nil {
        return nil, err
    }
    var persons []staff.PersonData
    for _, csv := range csvData {
        persons = append(persons, staff.PersonData{
            ID:            csv.ID,
            FirstName:     csv.FirstName,
            Surname:       csv.Surname,
            WorkEmail:     csv.WorkEmail,
            PersonalEmail: csv.PersonalEmail,
            CellPhone:     csv.CellPhone,
            Graduation:    csv.Graduation,
            Degree:        csv.Degree,
            LastPromotion: csv.LastPromotion,
            ActiveSince:   time.Now().String(),
        })
    }
    return persons, nil
}


func (r PersonCSV) ReadCSV(filename string) ([]PersonCSV, error) {
	options := ReadOptions{
		filename:  filename,
		delimiter: ';',
	}

	csvData, err := ReadData(options)

	if err != nil {
		return nil, err
	}
	return csvData, nil
}
