package util

import (
    "encoding/csv"
    "fmt"
    "os"
	"io"
	"golang.org/x/text/encoding/charmap"
	"bufio"
)

type PersonDTO struct {
    ID       string `json:"id" csv:"id"`
    Name     string `json:"name" csv:"name"`
    Age      int    `json:"age" csv:"age"`
    Location string `json:"location" csv:"location"`
}

func ReadData(filename string) (error) {
    // Open the CSV file
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error:", err)
        return err
    }
    defer file.Close()

    // Create a new reader that converts ISO-8859-1 to UTF-8
    reader := csv.NewReader(bufio.NewReader(charmap.ISO8859_1.NewDecoder().Reader(file)))

    // Read in the header row
    headers, err := reader.Read()
    if err != nil {
        fmt.Println("Error:", err)
        return err
    }

    // Print out the headers
    fmt.Println(headers)

 // Read in the records
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		fmt.Println(record)
	}

	return nil
}