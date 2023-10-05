package infra

import (
//	"IngSoftStaff/pkg/staff/domain"
	"fmt"
	//"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

  type PersonG struct {
	gorm.Model
	ID 		  uint
	FirstName     string
/*	Surname       string
	WorkEmail     string
	PersonalEmail string
	CellPhone     string
	Graduation    string
	Degree        string
	LastPromotion string
	ActiveSince   string
	*/
}

func InitDB() *gorm.DB {
	dbUrl := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=localtime"
	db, err := gorm.Open(postgres.Open(dbUrl),&gorm.Config{})
	//db, err := gorm.Open(sqlite.Open("../../../data/staff.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	  }

	  
	  db.Migrator().DropTable(&PersonG{})
	  if err != nil {
		panic("failed to drop table: " + err.Error())
	  	  }
		  
		 // Migrate the schema
	  //db.AutoMigrate(&PersonG{})
	  if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
	  return db

	  
	  
	  
	  
	  //p1 := domain.PersonData{ID: 1234, FirstName: "Gastón", Surname: "Mousques", WorkEmail: "mousques@academy.edu.uy", PersonalEmail: "mousques@example.com", CellPhone: "634323", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}

	  //db.Create(&p1)
}

func InsertPerson(db *gorm.DB) {
	  //p1 := domain.PersonData{ FirstName: "Gastón", Surname: "Mousques", WorkEmail: "mousques@academy.edu.uy", PersonalEmail: "mousques@example.com", CellPhone: "634323", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}

	  p1 := PersonG{ FirstName: "Gastón"}

	 result := db.Create(&p1)
	 fmt.Println(result.Error)

}



