package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ComputePractice2018/medicalcard/backend/data"
)

var testAppointment = "{\"date\":\"01.07.2018\",\"doc_name\":\"Иванов Пётр Данилович\",\"muse_name\":\"Прокофьева Вероника Максимовна\",\"complaint\":\"Чихает находясь в помещении с цветами\",\"checkup\":\"Покраснения слизистой\",\"diagnosis\":\"Стоматит\",\"treatment\":\"Изолироватть пациента от аллергена\"}"

func TestCrudHandlers(t *testing.T) {
	cl := data.NewAppointmentList()
	router := NewRouter(cl)

	req, err := http.NewRequest("GET", "/api/medicalcard/appointments", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 code")
	}
	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("Expected json MIME-type")
	}

	testData := strings.NewReader(testAppointment)
	req, err = http.NewRequest("POST", "/api/medicalcard/appointments", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Error("Expected 201 code")
	}

	if resp.Header.Get("Location") != "/api/medicalcard/appointments/0" {
		t.Error("Expected another location")
	}

	if len(cl.GetAppointments()) != 1 {
		t.Error("Expected new value")
	}

	testData = strings.NewReader(testAppointment)
	req, err = http.NewRequest("PUT", "/api/medicalcard/appointments/0", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Error("Expected 201 code")
	}

	if resp.Header.Get("Location") != "/api/medicalcard/appointments/0" {
		t.Error("Expected another location")
	}

	if len(cl.GetAppointments()) != 1 {
		t.Error("Expected old value")
	}

	req, err = http.NewRequest("DELETE", "/api/medicalcard/appointments/0", nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Error("Expected 204 code")
	}

	if len(cl.GetAppointments()) != 0 {
		t.Error("Expected old value")
	}
}
