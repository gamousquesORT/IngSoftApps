package main

import (
	"IngSoftStaff/pkg/staff/infra"
)

func main() {
	db:= infra.InitDB()

	infra.InsertPerson(db)
	
	
}


