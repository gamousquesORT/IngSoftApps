package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"IngSoftStaff/pkg/staff/domain"
)

  type PersonGORM struct {
	gorm.Model
	ID            uint32
	FirstName     string
	Surname       string
	WorkEmail     string
	PersonalEmail string
	CellPhone     string
	Graduation    string
	Degree        string
	LastPromotion string
	ActiveSince   string
}

func OpenDB() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=localtime"))
	if err != nil {
		panic("failed to connect database")
	  }

	  db.Migrator().DropTable(&PersonGORM{})
	  
	  // Migrate the schema
	  db.AutoMigrate(&PersonGORM{})


	  p1 := domain.PersonData{ID: 1234, FirstName: "Gastón", Surname: "Mousques", WorkEmail: "mousques@academy.edu.uy", PersonalEmail: "mousques@example.com", CellPhone: "634323", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}

	  db.Create(&p1)
}



func Init() *gorm.DB {
    dbURL := "postgres://postgres:postgres@localhost:5432/postgres"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Book{})

    return db
}