
package main

import (
	"fmt"
	"os"
	"time"
	"github.com/gocarina/gocsv"
)

type NotUsed struct {
	Name string
}

//https://articles.wesionary.team/easy-working-with-csv-in-golang-using-gocsv-package-9c8424728bbe


type Person struct {
	Id            string `csv:"-"`
	FirstName     string `csv:"Nombre"`
	Surname       string `csv:"Apellido"`
	WorkEmail     string `csv:"Email ORT"`
	PersonalEmail string `csv:"Email personal"`
	CellPhone     string `csv:"Celular"`
	Graduation    string `csv:"-"`
	Degree        string `csv:"-"`
	ActiveSince   time.Time `csv:"-"`
	Status        int `csv:"-"`
}

func main() {
	clientsFile, err := os.Open("../pkg/db/profesIngSoftOk.csv")
//	clientsFile, err := os.OpenFile("../pkg/db/profesIngSoftOk.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []*Person{}
	
	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}
	for _, client := range clients {
		fmt.Println("Hello", client.FirstName)
	}


}
