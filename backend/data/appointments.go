package data

import (
	"fmt"
)

// Appointment структура для хранения записи пациента
type Appointment struct {
	Date      string `json:"date"`
	DocName   string `json:"doc_name"`
	MuseName  string `json:"muse_name"`
	Complaint string `json:"complaint"`
	Checkup   string `json:"checkup"`
	Diagnosis string `json:"diagnosis"`
	Treatment string `json:"treatment"`
}

// appointments хранит список записей пациентов
var appointments []Appointment

// GetAppointments возвращает список приёмов
func GetAppointments() []Appointment {
	return appointments
}

// AddAppointment добавляем приём в конец списка и возвращает id
func AddAppointment(appointment Appointment) int {
	id := len(appointments)
	appointments = append(appointments, appointment)
	return id
}

// EditAppointment изменяет приём с id на appointment
func EditAppointment(appointment Appointment, id int) error {
	if id < 0 || id >= len(appointments) {
		return fmt.Errorf("Incorrect ID")
	}
	appointments[id] = appointment
	return nil
}

// RemoveAppointment удаляет приём по id
func RemoveAppointment(id int) error {
	if id < 0 || id >= len(appointments) {
		return fmt.Errorf("Incorrect ID")
	}
	appointments = append(appointments[:id], appointments[id+1:]...)
	return nil
}
