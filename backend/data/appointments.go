package data

import (
	"fmt"
)

// Appointment структура для хранения записи пациента
type Appointment struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	DocName   string `json:"doc_name"`
	MuseName  string `json:"muse_name"`
	Complaint string `json:"complaint"`
	Checkup   string `json:"checkup"`
	Diagnosis string `json:"diagnosis"`
	Treatment string `json:"treatment"`
}

// AppointmentList структура для списка записей пациентов
type AppointmentList struct {
	appointments []Appointment
}

// Editable интерфейс для работы со списком записей
type Editable interface {
	GetAppointments() []Appointment
	AddAppointment(appointment Appointment) int
	EditAppointment(appointment Appointment, id int) error
	RemoveAppointment(id int) error
}

// NewAppointmentList конструктор списка записей
func NewAppointmentList() *AppointmentList {
	return &AppointmentList{}
}

// GetAppointments возвращает список приёмов
func (cl *AppointmentList) GetAppointments() []Appointment {
	return cl.appointments
}

// AddAppointment добавляем приём в конец списка и возвращает id
func (cl *AppointmentList) AddAppointment(appointment Appointment) int {
	id := len(cl.appointments)
	appointment.ID = id
	cl.appointments = append(cl.appointments, appointment)
	return id
}

// EditAppointment изменяет приём с id на appointment
func (cl *AppointmentList) EditAppointment(appointment Appointment, id int) error {
	if id < 0 || id >= len(cl.appointments) {
		return fmt.Errorf("Incorrect ID")
	}
	cl.appointments[id] = appointment
	return nil
}

// RemoveAppointment удаляет приём по id
func (cl *AppointmentList) RemoveAppointment(id int) error {
	if id < 0 || id >= len(cl.appointments) {
		return fmt.Errorf("Incorrect ID")
	}
	cl.appointments = append(cl.appointments[:id], cl.appointments[id+1:]...)

	for i, appointment := range cl.appointments {
		appointment.ID = i
	}

	return nil
}
