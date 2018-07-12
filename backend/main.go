package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ComputePractice2018/medicalcard/backend/data"
	"github.com/ComputePractice2018/medicalcard/backend/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	connection := flag.String("connection", "medicalcard:SuperSecretPassword@tcp(db:3306)/medicalcard", "mysql connection string")
	flag.Parse()

	appointmentList, err := data.NewDBAppointmentList(*connection)
	defer appointmentList.Close()

	if err != nil {
		log.Fatalf("DB error: %+v", err)
	}

	router := server.NewRouter(appointmentList)
	log.Fatal(http.ListenAndServe(":8080", router))
}
