package main

import (
	"fmt"
	"time"
)

// Registro de estudiantes

type Student struct {
	Name              string
	Lastname          string
	DNI               string
	DateOfInscription time.Time
}

func (e Student) GetStudentDetail() {
	fmt.Printf("Name: %s\n Lastname: %s\n DNI: %s\n Date of inscription: %s",
		e.Name, e.Lastname, e.DNI, e.DateOfInscription)
}

func main() {
	student := Student{Name: "Ricardo", Lastname: "Gonzalez", DNI: "2420131022",
		DateOfInscription: time.Date(2013, 02, 04, 0, 0, 0, 0, time.UTC)}
	student.GetStudentDetail()
}
