package util

import (
	"IngSoftStaff/pkg/staff"
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {

	t.Run("given_a_file_returns_a_PersonList", func(t *testing.T) {
		want := []staff.Person{}
		options := ReadOptions{
			f:        CreateStaffList,
			filename: "../db/profesIngSoft1.csv",
			delimiter: ';',
		}

		got, err := ReadData(options)
		if err == nil {
			t.Errorf("Got not nil %v", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	
}
