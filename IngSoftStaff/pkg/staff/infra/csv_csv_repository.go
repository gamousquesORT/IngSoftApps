package infra

import (
	"IngSoftStaff/pkg/staff/domain"
	"encoding/csv"
	"io"
	"os"
	"regexp"
	"time"

	"github.com/gocarina/gocsv"
)

//https://articles.wesionary.team/easy-working-with-csv-in-golang-using-gocsv-package-9c8424728bbe

type NotUsed struct {
	Name string
}

type DateTime struct {
	time.Time
}

// Convert the internal date as CSV string
func (date *DateTime) MarshalCSV() (string, error) {
	return date.String(), nil
}

// You could also use the standard Stringer interface
func (date DateTime) String() string {
	return date.Time.Format("01022006")
}

// Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("01022006", csv)
	return err
}


type ReadOptions struct {
	filename string
	delimiter rune
}

type PersonCSV struct {
	ID            string `csv:"Persona"`
	FirstName     string `csv:"Nombre"`
	Surname       string `csv:"Apellido"`
	WorkEmail     string `csv:"Email_ORT"`
	PersonalEmail string `csv:"Email_Personal"`
	CellPhone     string `csv:"Celular"`
	Graduation    string `csv:"Graduado"`
	Degree        string `csv:"Grado"`
	LastPromotion string `csv:"FechaPromocion"`
}

type PersonCSVRepository struct {
	Options ReadOptions
	Filename string
}

func (p PersonCSV) ToPerson() *domain.PersonData {
	return &domain.PersonData{
		ID: p.ID,
		FirstName: p.FirstName,
		Surname: p.Surname,
		WorkEmail: p.WorkEmail,
		PersonalEmail: p.PersonalEmail,
		CellPhone: p.CellPhone,
		Graduation: p.Graduation,
		Degree: p.Degree,
	}

}
var _ domain.PersonRepository = &PersonCSVRepository{}

func NewPersonCSVRepository(filename string) domain.PersonRepository {
	return &PersonCSVRepository{
		Options:  ReadOptions{filename: filename, delimiter: ','},
		Filename: filename,
	}
}

func (repo PersonCSVRepository) ImportData() (*[]domain.PersonData, error) {

	 delimiter, err := readCSVDelimiter(repo.Options.filename, repo.Options.delimiter);
	 if err != nil {
		return nil, err
	}

	//override default csv reader delimiter
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = delimiter
		return r
	})

	staffFile, err := os.Open(repo.Options.filename)
	if err != nil {
		return nil, err
	}
	defer staffFile.Close()

	personsInCSVFormat := []PersonCSV{}

	if err := gocsv.UnmarshalFile(staffFile, &personsInCSVFormat); err != nil { // Load clients from file
		return nil, err
	}

	personsInDomainDataFormat := []domain.PersonData{}
	for _, p := range personsInCSVFormat {
		personsInDomainDataFormat = append(personsInDomainDataFormat, *p.ToPerson())
	}

	return &personsInDomainDataFormat, nil

}


func readCSVDelimiter(fileName string, delimiter rune) (rune, error) {
	file, err := os.Open(fileName)
    if err != nil {
        return 0, err
    }
    defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = delimiter
	csvReader.FieldsPerRecord = -1

	// read csv header
	var header []string
	header, err = csvReader.Read()
	if err != nil {
        return 0, err
    }

	var r rune = csvReader.Comma
	var ok bool

	r, ok = findSeparator(header)
	if !ok {
		return 0, nil
	}
	
	return r, nil

}


func findSeparator(s []string) (rune, bool) {

	var match string = ","

	if len(s) == 1 {
		re := regexp.MustCompile(`[;,]`)
    	match = re.FindString(s[0])
		if match == "" {
			return 0, false
		} 
	}
    return rune(match[0]), true
}