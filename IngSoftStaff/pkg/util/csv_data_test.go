package util

import (
	"IngSoftStaff/pkg/staff"
	"fmt"
	"reflect"
	"testing"
)



func TestReadData(t *testing.T) {

	t.Run("given_a_Non_Existent_file_returns_an_Error", func(t *testing.T) {
		got := []staff.Person{}
	

		options := ReadOptions{
			filename: "../../../data/profesIngSoft1.csv",
			delimiter: ';',
		}
		want := []staff.Person{}
		newgot, err  := ReadData(options, got)
		if err != nil {
			t.Fatalf("Got not nil %v", err.Error())
		}

		if !reflect.DeepEqual(newgot, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("given_a_wellformed_file_returns_a_PersonList", func(t *testing.T) {
		got := []staff.Person{}
	

		options := ReadOptions{
			//CASA filename: "../../../data/profesIngSoft.csv",
			filename: "../../data/profesIngSoft.csv",
			delimiter: ',',
		}

		
		p1:= staff.Person{Id: "1234", FirstName: "Gastón", Surname: "Mousques", WorkEmail: "mousques@academy.edu.uy", PersonalEmail: "mousques@example.com", CellPhone: "634323", Graduation: "SI", Degree: ""}
		p2:= staff.Person{Id: "14", FirstName: "Ignacio", Surname: "Valla", WorkEmail: "ignacio.valla@academy.edu.uy", PersonalEmail: "ignacio.valla@example.com", CellPhone: "694870", Graduation: "SI", Degree: ""}
		p3 := staff.Person{Id: "9526", FirstName: "Martin", Surname: "Sol", WorkEmail: "solar_m@academy.edu.uy", PersonalEmail: "", CellPhone: "2312255", Graduation: "SI", Degree: ""}
		p4 := staff.Person{Id: "404420", FirstName: "Alejandro", Surname: "Ado", WorkEmail: "ado@academy.edu.uy", PersonalEmail: "Alejandro@example.com", CellPhone: "6761532", Graduation: "SI", Degree: ""}
		p5 := staff.Person{Id: "213518", FirstName: "Valeria", Surname: "Ferreiro", WorkEmail: "valeria.ferreiro@academy.edu.uy", PersonalEmail: "valferper@example.com", CellPhone: "66363844", Graduation: "NO", Degree: ""}
		p6 := staff.Person{Id: "134058", FirstName: "Patricia", Surname: "De León", WorkEmail: "pd134058@academy.edu.uy", PersonalEmail: "patricia.deleon@example.com", CellPhone: "66400505", Graduation: "SI", Degree: ""}
		p7 := staff.Person{Id: "214157", FirstName: "Roxana", Surname: "Falco", WorkEmail: "rf214157@academy.edu.uy", PersonalEmail: "roxana.falco20@example.com", CellPhone: "6584563", Graduation: "SI", Degree: ""}
		p8 := staff.Person{Id: "228175", FirstName: "Johny", Surname: "Kidd", WorkEmail: "jk228175@academy.edu.uy", PersonalEmail: "kidnixt@example.com", CellPhone: "6385172", Graduation: "NO", Degree: ""}
		p9 := staff.Person{Id: "126376", FirstName: "Juan Pablo", Surname: "Russo", WorkEmail: "juan.russo@academy.edu.uy", PersonalEmail: "jprussoibanez@example.com", CellPhone: "476500", Graduation: "SI", Degree: ""}
		p10 := staff.Person{Id: "164538", FirstName: "Analía", Surname: "Moreira", WorkEmail: "analia.moreira@academy.edu.uy", PersonalEmail: "analiasmoreira@example.com", CellPhone: "6120290", Graduation: "SI", Degree: ""}
		p11 := staff.Person{Id: "157727", FirstName: "Fabiana", Surname: "Pedrini", WorkEmail: "fp157727@academy.edu.uy", PersonalEmail: "fpedrini19@example.com", CellPhone: "91211826", Graduation: "SI", Degree: ""}
		p12 := staff.Person{Id: "203067", FirstName: "Alejo", Surname: "Pereira", WorkEmail: "alejo.pereira@academy.edu.uy", PersonalEmail: "alejo.pereira13@example.com", CellPhone: "6750461", Graduation: "SI", Degree: ""}
		p13 := staff.Person{Id: "61505", FirstName: "Juan", Surname: "Ort", WorkEmail: "Juan.ort@academy.edu.uy", PersonalEmail: "Juan.ort1@example.com", CellPhone: "5641685", Graduation: "SI", Degree: ""}
		p14 :=staff.Person{Id: "149128", FirstName: "Carolina", Surname: "Fontón", WorkEmail: "caro.fontan@academy.edu.uy", PersonalEmail: "cp@example.com", CellPhone: "600637", Graduation: "SI", Degree: ""}

		want := []staff.Person{p1,p2,p3,p4,p5,p6,p7,p8,p9,p10,p11,p12,p13,p14}
		
		newgot, err  := ReadData(options, got)
		if err != nil {
			t.Errorf("Got not nil %v", err)
		}

		/*
		if !reflect.DeepEqual(newgot, want) {
			t.Errorf("got %v want %v", newgot, want)
		}
		*/

		for i := range newgot {
			if newgot[i].Id != want[i].Id {
				fmt.Printf("** [%v] got %v want %v \n", i, newgot[i].Id, want[i].Id)
			}
			if newgot[i].FirstName != want[i].FirstName {	
				fmt.Printf("** [%v] got %v want %v  \n", i, newgot[i].FirstName, want[i].FirstName)
			}	
			if newgot[i].Surname != want[i].Surname {
				fmt.Printf("** [%v] got %v want %v \n", i, newgot[i].Surname, want[i].Surname)
			}

			if newgot[i].WorkEmail != want[i].WorkEmail {
				fmt.Printf("** [%v] got %v want %v \n", i, newgot[i].WorkEmail, want[i].WorkEmail)
			}

			if newgot[i].PersonalEmail != want[i].PersonalEmail {
				fmt.Printf("** [%v] got %v want %v \n", i, newgot[i].PersonalEmail, want[i].PersonalEmail)
			}

			if newgot[i].CellPhone != want[i].CellPhone {
				fmt.Printf("** [%v] got %v want %v \n", i, newgot[i].CellPhone, want[i].CellPhone)
			}

			if newgot[i].Graduation != want[i].Graduation {
				fmt.Printf("** [%v] got %v want %v \n", i, newgot[i].Graduation, want[i].Graduation)
			}

			if newgot[i].Degree != want[i].Degree {
				fmt.Printf("** [%v] got %v want %v \n", i, newgot[i].Degree, want[i].Degree)
			}
			fmt.Printf("--------------------------------------------------\n")
			
		}

	})


}

