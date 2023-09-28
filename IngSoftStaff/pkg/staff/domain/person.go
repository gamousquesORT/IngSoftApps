package domain

import (
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

type PersonRepository interface {
	ImportData() (*[]PersonData, error)
}






