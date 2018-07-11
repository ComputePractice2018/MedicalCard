package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ComputePractice2018/medicalcard/backend/data"
	"github.com/gorilla/mux"
)

// GetAppointments обрабатывает запросы на получение списка записей пациентов
func GetAppointments(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(cl.GetAppointments())
		if err != nil {
			message := fmt.Sprintf("Unable to encode data: %v", err)
			http.Error(w, message, http.StatusInternalServerError)
			log.Println(message)
			return
		}
	}
}

// AddAppointment обрабатывает POST запрос
func AddAppointment(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var appointment data.Appointment
		err := json.NewDecoder(r.Body).Decode(&appointment)
		if err != nil {
			message := fmt.Sprintf("Unable to decode POST data: %v", err)
			http.Error(w, message, http.StatusUnsupportedMediaType)
			log.Println(message)
			return
		}
		id := cl.AddAppointment(appointment)
		w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
		w.WriteHeader(http.StatusCreated)
	}
}

// EditAppointment ообрабатывает PUT запрос
func EditAppointment(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var appointment data.Appointment
		err := json.NewDecoder(r.Body).Decode(&appointment)
		if err != nil {
			message := fmt.Sprintf("Unable to decode PUT data: %v", err)
			http.Error(w, message, http.StatusUnsupportedMediaType)
			log.Println(message)
			return
		}
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}

		err = cl.EditAppointment(appointment, id)
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusAccepted)
	}
}

// DeleteAppointment обрабатывает DELETE запрос
func DeleteAppointment(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}

		err = cl.RemoveAppointment(id)
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusNoContent)
	}
}
