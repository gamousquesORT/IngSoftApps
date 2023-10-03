package infra

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"IngSoftStaff/pkg/staff/domain"
)

  type PersonGORM struct {
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
	dbUrl := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=localtime"
	db, err := gorm.Open(postgres.Open(dbUrl),&gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	  }

	  //db.Migrator().DropTable(&PersonGORM{})
	  
	  // Migrate the schema
	  //db.AutoMigrate(&PersonGORM{})


	  p1 := domain.PersonData{ID: 1234, FirstName: "Gastón", Surname: "Mousques", WorkEmail: "mousques@academy.edu.uy", PersonalEmail: "mousques@example.com", CellPhone: "634323", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}

	  db.Create(&p1)
}



