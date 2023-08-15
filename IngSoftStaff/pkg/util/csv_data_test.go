package util

import (
	"IngSoftStaff/mocks"
	"IngSoftStaff/pkg/staff"
	"reflect"
	"testing"
    "time"
)

func TestReadData(t *testing.T) {
	// Test with a valid file
	var reader mocks.MockCSVReader
	reader.Init()
    
	t.Run("given_empty_file_returns_emptyDTO", func(t *testing.T) {
		want := []staff.PersonDTO{}
		got, err := ReadData("../db/profesIngSoft1.csv")

		if err == nil {
			t.Errorf("Got not nil %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("given_a_file_Reads_DTOs", func(t *testing.T) {

		want := []staff.PersonDTO{
            {
                IDPersona:     "1",
                FirstName:     "Alice",
                Surname:       "123",
                WorkEmail:     "alice@example.com",
                PersonalEmail: "alice@gmail.com",
                CellPhone:     "555-1234",
                Graduation:    time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
                Degree:        1,
            },
            {
                IDPersona:     "2",
                FirstName:     "Bob",
                Surname:       "456",
                WorkEmail:     "bob@example.com",
                PersonalEmail: "bob@gmail.com",
                CellPhone:     "555-5678",
                Graduation:    time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
                Degree:        2,
            },
        }
		got, err := ReadData("../db/profesIngSoft.csv")

		if err != nil {
			t.Errorf("Expected non-nil error, but got nil")
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
