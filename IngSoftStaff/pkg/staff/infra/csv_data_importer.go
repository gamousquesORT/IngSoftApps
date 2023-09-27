package staff

import (
	"IngSoftStaff/pkg/staff"
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


type ReadOptions struct {
	filename string
	delimiter rune
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
	staff []PersonCSV
	options ReadOptions
}

func (p PersonCSV) ToPerson() *staff.PersonData {
	return &staff.PersonData{
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
var _ staff.PersonRepository = &PersonCSVRepository{}

func NewPersonCSVRepository(filename string) staff.PersonRepository {
	return &PersonCSVRepository{
		staff: []PersonCSV{},
		options: ReadOptions{
			filename: filename,
			delimiter: ';',
		},
	}
}

func (repo PersonCSVRepository) ImportData() (*[]staff.PersonData, error) {

	 delimiter, err := ReadCSVHeader(repo.options.filename);
	 if err != nil {
		return nil, err
	}

	// set the pipe as the delimiter for reading
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = delimiter
		return r
	})

	staffFile, err := os.Open(repo.options.filename)

	if err != nil {
		return nil, err
	}
	defer staffFile.Close()


	if err := gocsv.UnmarshalFile(staffFile, &repo.staff); err != nil { // Load clients from file
		return nil, err
	}

	persons := []staff.PersonData{}

	for _, p := range repo.staff {
		persons = append(persons, *p.ToPerson())
	}

	return &persons, nil

}


func ReadCSVHeader(fileName string) (rune, error) {
	file, err := os.Open(fileName)
    if err != nil {
        return 0, err
    }

    // remember to close the file at the end of the program
    defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ','
	csvReader.FieldsPerRecord = -1

	// read csv header
	var header []string
	header, err = csvReader.Read()
	if err != nil {
        return 0, err
    }

	var r rune = csvReader.Comma
	var ok bool

	if len(header) == 1 {
		r, ok = FindSeparator(header[0])
		if !ok {
			return 0, nil
		}
	}

	return r, nil

}


func FindSeparator(s string) (rune, bool) {
    re := regexp.MustCompile(`[;,]`)
    match := re.FindString(s)
    if match == "" {
        return 0, false
    }
    return rune(match[0]), true
}