package util

import (
    "os"
    "io"
    "IngSoftStaff/pkg/staff"
    "encoding/csv"
)

func ReadData(filename string) ([]staff.PersonDTO, error) {

    var persons []staff.PersonDTO

    records := readCSV(filename)

    return persons, nil
}
func readCSV(filename string) ([][]string, error) {

    records := [][]string{};

    file, err := os.Open(filename)
    if err != nil {
        return [][]string{}, err
    }
    defer file.Close()
        
    reader := csv.NewReader(file)
        
    // read csv values using csv.Reader
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return []staff.PersonDTO{}, err
        }
        // do something with read line
        records = append(records, record)

        return records, nil
    }
}