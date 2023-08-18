package util

import (
	"IngSoftStaff/pkg/staff"
	"encoding/csv"
	"os"
)

// try using https://github.com/gocarina/gocsv
// check https://medium.com/@KeithAlpichi/gos-standard-library-by-example-encoding-csv-75f098169822

type ReadOptions struct {
	f func(data [][]string) (staffList []staff.Person)
	filename string
	delimiter rune
}

func ReadData(opt ReadOptions) ([]staff.Person, error) {
	file, err := os.Open(opt.filename)
    if err != nil {
        return nil, err
    }

    // remember to close the file at the end of the program
    defer file.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(file)
	csvReader.Comma = opt.delimiter
    data, err := csvReader.ReadAll()
    if err != nil {
        return nil, err
    }

    // convert records to array of structs
    staffList := opt.f(data)
	return staffList, nil
}

func CreateStaffList(data [][]string) (staffList []staff.Person) {
	staffList = []staff.Person{}
	for _, record := range data {
		println(record)
	}
	return staffList
}


