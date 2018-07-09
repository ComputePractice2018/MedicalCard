package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/medicalcard/backend/data"
)

// AppointmentsHandler обрабатывает все запросы к /api/medicalcard/appointment
func AppointmentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetAppointments(w, r)
		return
	}

	if r.Method == "POST" {
		AddAppointment(w, r)
		return
	}

	message := fmt.Sprintf("Method %s is not allowed", r.Method)
	http.Error(w, message, http.StatusMethodNotAllowed)
	log.Println(message)
}

// GetAppointments обрабатывает запросы на получение списка записей пациентов
func GetAppointments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request method: %s", r.Method)
	binaryData, err := json.Marshal(data.AppointmentList)
	if err != nil {
		message := fmt.Sprintf("JSON marshaling error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(binaryData)
	if err != nil {
		message := fmt.Sprintf("Handler write error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
	}
}

// AddAppointment обрабатывает POST запрос
func AddAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment data.Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		message := fmt.Sprintf("Unable to decode POST data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	log.Printf("%+v", appointment)
	w.WriteHeader(http.StatusCreated)
}
