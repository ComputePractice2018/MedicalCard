package server

import (
	"github.com/ComputePractice2018/medicalcard/backend/data"
	"github.com/gorilla/mux"
)

func NewRouter(appointmentList data.Editable) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/medicalcard/appointments", GetAppointments(appointmentList)).Methods("GET")
	router.HandleFunc("/api/medicalcard/appointments", AddAppointment(appointmentList)).Methods("POST")
	router.HandleFunc("/api/medicalcard/appointments/{id}", EditAppointment(appointmentList)).Methods("PUT")
	router.HandleFunc("/api/medicalcard/appointments/{id}", DeleteAppointment(appointmentList)).Methods("DELETE")
	return router
}
