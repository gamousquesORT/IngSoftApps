package util

import (
	"IngSoftStaff/pkg/staff"
	"reflect"
	"testing"
)

func TestReadData(t *testing.T) {

	t.Run("given_a_Non_Existent_file_returns_an_Error", func(t *testing.T) {
		got := []staff.Person{}
	

		options := ReadOptions{
			filename: "../../profesIngSoft1.csv",
			delimiter: ';',
		}
		want := []staff.Person{}
		newgot, err  := ReadData(options, got)
		if err != nil {
			t.Fatalf("Got not nil %v", err.Error())
		}

		if !reflect.DeepEqual(newgot, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("given_a_wellformed_file_returns_a_PersonList", func(t *testing.T) {
		got := []staff.Person{}
	

		options := ReadOptions{
			filename: "../../profesIngSoft.csv",
			delimiter: ';',
		}
		want := []staff.Person{}
		newgot, err  := ReadData(options, got)
		if err != nil {
			t.Errorf("Got not nil %v", err)
		}

		if !reflect.DeepEqual(newgot, want) {
			t.Errorf("got %v want %v", newgot, want)
		}
	})


	func (t staff.Person) String() string {
		return fmt.Sprintf("{%v %v }", t.Id, t.FirstName)
	}
}
