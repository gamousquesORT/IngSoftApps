package tests

import (
	//"fmt"
	//"reflect"
	"testing"

	"IngSoftStaff/pkg/staff/domain"
	"IngSoftStaff/pkg/staff/infra"
)

func checkCSV(newgot *[]domain.PersonData, want []domain.PersonData) bool {
	var cnt int = 0	

	for i := range *newgot {
		if (*newgot)[i].ID != want[i].ID {
			cnt++
		}
		if (*newgot)[i].FirstName != want[i].FirstName {
			cnt++
		}
		if (*newgot)[i].Surname != want[i].Surname {
			cnt++
		}

		if (*newgot)[i].WorkEmail != want[i].WorkEmail {
			cnt++
		}

		if (*newgot)[i].PersonalEmail != want[i].PersonalEmail {
			cnt++
		}

		if (*newgot)[i].CellPhone != want[i].CellPhone {
			cnt++
		}

		if (*newgot)[i].Graduation != want[i].Graduation {
			cnt++
		}

		if (*newgot)[i].Degree != want[i].Degree {
			cnt++
		}

	 }
	return cnt == 0
}


func TestReadData(t *testing.T) {

	t.Run("given_a_Non_Existent_file_returns_an_Error", func(t *testing.T) {


		repo := infra.NewPersonCSVRepository("../../../data/profesIngSoft22.csv")
		_, err := repo.ImportData();
	
		if err == nil {
			t.Fatalf("Got not nil %v", err.Error())
		}

	})

	t.Run("given_a_wellformed_file_returns_a_PersonList", func(t *testing.T) {

		p1 := domain.PersonData{ID: "1234", FirstName: "Gastón", Surname: "Mousques", WorkEmail: "mousques@academy.edu.uy", PersonalEmail: "mousques@example.com", CellPhone: "634323", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}
		p2 := domain.PersonData{ID: "14", FirstName: "Ignacio", Surname: "Valla", WorkEmail: "ignacio.valla@academy.edu.uy", PersonalEmail: "ignacio.valla@example.com", CellPhone: "694870", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}
		p3 := domain.PersonData{ID: "9526", FirstName: "Martin", Surname: "Sol", WorkEmail: "solar_m@academy.edu.uy", PersonalEmail: "", CellPhone: "2312255", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}
		p4 := domain.PersonData{ID: "404420", FirstName: "Alejandro", Surname: "Ado", WorkEmail: "ado@academy.edu.uy", PersonalEmail: "Alejandro@example.com", CellPhone: "6761532", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}
		p5 := domain.PersonData{ID: "213518", FirstName: "Valeria", Surname: "Ferreiro", WorkEmail: "valeria.ferreiro@academy.edu.uy", PersonalEmail: "valferper@example.com", CellPhone: "66363844", Graduation: "NO", Degree: "", LastPromotion: "23/09/2023"}
		p6 := domain.PersonData{ID: "134058", FirstName: "Patricia", Surname: "De León", WorkEmail: "pd134058@academy.edu.uy", PersonalEmail: "patricia.deleon@example.com", CellPhone: "66400505", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}
		p7 := domain.PersonData{ID: "214157", FirstName: "Roxana", Surname: "Falco", WorkEmail: "rf214157@academy.edu.uy", PersonalEmail: "roxana.falco20@example.com", CellPhone: "6584563", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}
		p8 := domain.PersonData{ID: "228175", FirstName: "Johny", Surname: "Kidd", WorkEmail: "jk228175@academy.edu.uy", PersonalEmail: "kidnixt@example.com", CellPhone: "6385172", Graduation: "NO", Degree: "", LastPromotion: "23/09/2023"}
		p9 := domain.PersonData{ID: "126376", FirstName: "Juan Pablo", Surname: "Russo", WorkEmail: "juan.russo@academy.edu.uy", PersonalEmail: "jprussoibanez@example.com", CellPhone: "476500", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}
		p10 := domain.PersonData{ID: "164538", FirstName: "Analía", Surname: "Moreira", WorkEmail: "analia.moreira@academy.edu.uy", PersonalEmail: "analiasmoreira@example.com", CellPhone: "6120290", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}
		p11 := domain.PersonData{ID: "157727", FirstName: "Fabiana", Surname: "Pedrini", WorkEmail: "fp157727@academy.edu.uy", PersonalEmail: "fpedrini19@example.com", CellPhone: "91211826", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}
		p12 := domain.PersonData{ID: "203067", FirstName: "Alejo", Surname: "Pereira", WorkEmail: "alejo.pereira@academy.edu.uy", PersonalEmail: "alejo.pereira13@example.com", CellPhone: "6750461", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}
		p13 := domain.PersonData{ID: "61505", FirstName: "Juan", Surname: "Ort", WorkEmail: "Juan.ort@academy.edu.uy", PersonalEmail: "Juan.ort1@example.com", CellPhone: "5641685", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}
		p14 := domain.PersonData{ID: "149128", FirstName: "Carolina", Surname: "Fontón", WorkEmail: "caro.fontan@academy.edu.uy", PersonalEmail: "cp@example.com", CellPhone: "600637", Graduation: "SI", Degree: "", LastPromotion: "23/09/2023"}

		want := []domain.PersonData{p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14}

		var filenameCASA = "../../../data/profesIngSoft.csv"
		//filenameORT:  "../../data/profesIngSoft.csv"
		repo := infra.NewPersonCSVRepository(filenameCASA)
		got, err := repo.ImportData();
	

		if err != nil {
			t.Errorf("Got not nil %v", err)
		}

		if !checkCSV(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

}
