package data

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

// AppointmentList хранит список записей пациентов
var AppointmentList = []Appointment{Appointment{
	Date:      "01.07.2018",
	DocName:   "Смирнов Василий Данилович",
	MuseName:  "Прокофьева Вероника Максимовна",
	Complaint: "Чихает находясь в помещении с цветами",
	Checkup:   "Покраснения слизистой",
	Diagnosis: "Стоматит",
	Treatment: "Изолироватть пациента от аллергена",
}}
