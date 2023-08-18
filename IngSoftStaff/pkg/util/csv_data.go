package util

import (
	"IngSoftStaff/pkg/staff"
	"encoding/csv"
	"os"
	"io"
	"time"
)

// try using https://github.com/gocarina/gocsv
// check https://medium.com/@KeithAlpichi/gos-standard-library-by-example-encoding-csv-75f098169822

type ReadOptions struct {
	f func(data []string, staffList []staff.Person) ([]staff.Person, error)
	filename string
	delimiter rune
}

func ReadData(opt ReadOptions,staffList []staff.Person) ([]staff.Person, error) {
	file, err := os.Open(opt.filename)
    if err != nil {
        return nil, err
    }

    // remember to close the file at the end of the program
    defer file.Close()


	csvReader := csv.NewReader(file)
	csvReader.Comma = opt.delimiter
	csvReader.FieldsPerRecord = -1

	// read csv header

	_, err = csvReader.Read()
	if err != nil {
        return nil, err
    }

	var newStaffList []staff.Person

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
		  break
		} else if err != nil {
		  panic(err) // or handle it another way
		}
	
		newStaffList, err = opt.f(row, staffList)
		if err != nil {
			return nil, err
		}
		staffList = newStaffList
	  }

  
	return newStaffList, nil
}

func AddToStaffList(data []string, staffList []staff.Person) ([]staff.Person, error) {

	timeStr := "2021-08-25T12:34:56Z"
	layout := "2006-01-02T15:04:05Z"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return nil, err
	}

	newPerson := staff.NewPerson(data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7],
		t, 1)

	println(data)

	staffList = append(staffList, newPerson)
	
	return staffList, nil
}


