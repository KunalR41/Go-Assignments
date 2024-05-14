package main

import (
	"fmt"
)

// Chain_of_Rasponseblity
type Patient struct {
	Name     string
	Age      int
	Severity string
}

type Handler interface {
	HandlePatient(patient *Patient)
	SetNext(handler Handler)
}

type BaseHandler struct {
	nextHandler Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

type ReceptionHandler struct {
	BaseHandler
}

func (h *ReceptionHandler) HandlePatient(patient *Patient) {
	fmt.Printf("Patient %s registered at reception.\n", patient.Name)
	if h.nextHandler != nil {
		h.nextHandler.HandlePatient(patient)
	}
}

type DoctorHandler struct {
	BaseHandler
}

func (h *DoctorHandler) HandlePatient(patient *Patient) {
	if patient.Severity == "Emergency" {
		fmt.Printf("Patient %s is critical. Immediate attention required!\n", patient.Name)
	} else if patient.Severity == "Routine" {
		fmt.Printf("Patient %s is assigned to a general practitioner.\n", patient.Name)
	} else {
		fmt.Printf("Patient %s is assigned to a specialist.\n", patient.Name)
	}
}

func main() {

	receptionHandler := &ReceptionHandler{}
	doctorHandler := &DoctorHandler{}

	receptionHandler.SetNext(doctorHandler)

	patient1 := &Patient{Name: "Jay", Age: 35, Severity: "Routine"}
	patient2 := &Patient{Name: "Ram", Age: 45, Severity: "Emergency"}
	patient3 := &Patient{Name: "Sham", Age: 25, Severity: "Specialist"}

	receptionHandler.HandlePatient(patient1)
	receptionHandler.HandlePatient(patient2)
	receptionHandler.HandlePatient(patient3)
}
