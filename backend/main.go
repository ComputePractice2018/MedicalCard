package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/medicalcard/backend/server"
)

func main() {
	//a := flag.Int("a", 1, "число 1")
	//b := flag.Int("b", 2, "число 2")
	//flag.Parse()

	http.HandleFunc("/api/medicalcard/appointment", server.GetAppointments)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
