package util

import (
	"IngSoftStaff/pkg/staff"
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {

	t.Run("given_a_wellformedfile_returns_a_PersonList", func(t *testing.T) {
		got := []staff.Person{}
	

		options := ReadOptions{
			f:        AddToStaffList,
			filename: "../db/profesIngSoft1.csv",
			delimiter: ';',
		}
		want := []staff.Person{}
		newgot, err  := ReadData(options, got)
		if err != nil {
			t.Errorf("Got not nil %v", err)
		}

		if !reflect.DeepEqual(newgot, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	
}
