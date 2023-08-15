package mocks


import (
    "io"
    "IngSoftStaff/pkg/staff"
)

type CSVReader interface {
    Read() ([]string, error)
    ReadAll() ([][]string, error)
    Init()
}

type RecordMetadata struct {
    Field1 string
    Field2 string
    Field3 string
    Field4 string
    Field5 string
    Field6 string
    Field7 string
    Field8 string
}

type MockCSVReader struct {
    File    string
    Records []staff.PersonDTO
    Index   int
}

func (m *MockCSVReader) Init() {
    m.File = "../db/profesIngSoft.csv"

    m.Records = []staff.PersonDTO{
            {
                IDPersona: "IDPersona",
                FirstName: "FirstName",
                Surname: "Surname",
                WorkEmail: "WorkEmail",
                PersonalEmail: "PersonalEmail",
                CellPhone: "CellPhone",
                Graduation: "Graduation",
                Degree: "Degree",
            },
            {
                IDPersona:     "1",
                FirstName:     "Alice",
                Surname:       "123",
                WorkEmail:     "alice@example.com",
                PersonalEmail: "alice@gmail.com",
                CellPhone:     "555-1234",
                Graduation:    "2006-01-02",
                Degree:        "1",
            },
            {
                IDPersona:     "2",
                FirstName:     "Bob",
                Surname:       "456",
                WorkEmail:     "bob@example.com",
                PersonalEmail: "bob@gmail.com",
                CellPhone:     "555-5678",
                Graduation:    "2016-01-02",
                Degree:        "2",
            },
            {
                IDPersona:     "3",
                FirstName:     "Charlie",
                Surname:       "789",
                WorkEmail:     "charlie@example.com",
                PersonalEmail: "charlie@gmail.com",
                CellPhone:     "555-9012",
                Graduation:    "2016-01-02",
                Degree:       "3",
            },
            {
                IDPersona:     "4",
                FirstName:     "Dave",
                Surname:       "101112",
                WorkEmail:     "dave@example.com",
                PersonalEmail: "dave@gmail.com",
                CellPhone:     "555-1314",
                Graduation:    "2018-01-02",
                Degree:        "4",
            },
            {
                IDPersona:     "5",
                FirstName:     "Eve",
                Surname:       "151617",
                WorkEmail:     "eve@example.com",
                PersonalEmail: "eve@gmail.com",
                CellPhone:     "555-1819",
                Graduation:    "1986-01-02",
                Degree:        "5",
            },
    }

    m.Index = 0
}

func (m *MockCSVReader) Read() ([]string, error) {
    if m.File != "../db/profesIngSoft.csv" {
        return []string, io.EOF
    }
    if m.Index >= len(m.Records) {
        return nil, io.EOF
    }

    record := m.Records[m.Index]
    m.Index++
    return record, nil
}