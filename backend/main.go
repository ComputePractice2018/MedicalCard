package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ComputePractice2018/medicalcard/backend/data"
	"github.com/ComputePractice2018/medicalcard/backend/server"
)

func main() {
	//a := flag.Int("a", 1, "число 1")
	//b := flag.Int("b", 2, "число 2")
	//flag.Parse()
	appointmentList := data.NewAppointmentList()
	router := mux.NewRouter()
	router.HandleFunc("/api/medicalcard/appointments", server.GetAppointments(appointmentList)).Methods("GET")
	router.HandleFunc("/api/medicalcard/appointments", server.AddAppointment(appointmentList)).Methods("POST")
	router.HandleFunc("/api/medicalcard/appointments/{id}", server.EditAppointment(appointmentList)).Methods("PUT")
	router.HandleFunc("/api/medicalcard/appointments/{id}", server.DeleteAppointment(appointmentList)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
