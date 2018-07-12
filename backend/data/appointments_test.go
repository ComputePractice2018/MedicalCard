package data

import "testing"

var testAppointments = []Appointment{
	{
		ID:        1,
		Date:      "10.11.2005",
		DocName:   "Иванов Иван Иванович",
		MuseName:  "Петров Пётр Петрович",
		Complaint: "Кашляет",
		Checkup:   "Кашель",
		Diagnosis: "Грипп",
		Treatment: "Таблетки",
	},
	{
		ID:        2,
		Date:      "07.07.2007",
		DocName:   "Сидоров Сидр Сидорович",
		MuseName:  "Малахов Андрей Васильевич",
		Complaint: "Чихает",
		Checkup:   "Посраснения слизистой",
		Diagnosis: "Аллергия",
		Treatment: "Уколы",
	},
}

func TestAddAppointment(t *testing.T) {
	cl := NewAppointmentList()

	cl.AddAppointment(testAppointments[0])

	if cl.GetAppointments()[0] != testAppointments[0] {
		t.Errorf("AddAppointment in not working")
	}
}

func TestEditAppointment(t *testing.T) {
	cl := NewAppointmentList()
	cl.AddAppointment(testAppointments[0])

	err := cl.EditAppointment(testAppointments[1], 1)

	if cl.GetAppointments()[0] != testAppointments[1] {
		t.Errorf("EditAppointment in not working")
	}
	if err != nil {
		t.Errorf("Unexpected EditAppointment error: %+v", err)
	}

	err = cl.EditAppointment(testAppointments[1], -1)
	if err == nil {
		t.Errorf("Nothandled out of range error")
	}
}

func TestRemoveAppointment(t *testing.T) {
	cl := NewAppointmentList()
	cl.AddAppointment(testAppointments[0])
	cl.AddAppointment(testAppointments[1])

	err := cl.RemoveAppointment(1)

	if cl.GetAppointments()[0] != testAppointments[1] {
		t.Errorf("RemoveAppointment in not working")
	}
	if err != nil {
		t.Errorf("Unexpected RemoveAppointment error")
	}

	err = cl.RemoveAppointment(-1)
	if err == nil {
		t.Errorf("Nothandled out of range error")
	}
}
