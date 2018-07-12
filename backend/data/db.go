package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type DBAppointmentList struct {
	db *gorm.DB
}

func NewDBAppointmentList(connection string) (*DBAppointmentList, error) {
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Appointment{})
	return &DBAppointmentList{db: db}, db.Error
}

func (cl *DBAppointmentList) Close() {
	cl.db.Close()
}

func (cl *DBAppointmentList) GetAppointments() []Appointment {
	var appointments []Appointment
	cl.db.Find(&appointments)
	return appointments
}

func (cl *DBAppointmentList) AddAppointment(appointment Appointment) int {
	cl.db.Create(&appointment)
	return appointment.ID
}

func (cl *DBAppointmentList) EditAppointment(appointment Appointment, id int) error {
	var appointments []Appointment
	cl.db.Where("id = ?", id).Find(&appointments)
	if len(appointments) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}

	appointment.ID = appointments[0].ID
	cl.db.Save(&appointment)
	return cl.db.Error
}

func (cl *DBAppointmentList) RemoveAppointment(id int) error {
	var appointments []Appointment
	cl.db.Where("id = ?", id).Find(&appointments)
	if len(appointments) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}
	cl.db.Delete(&appointments[0])
	return cl.db.Error
}
