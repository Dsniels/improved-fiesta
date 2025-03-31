package entities

import "fmt"

type Profesor struct {
	Id       int
	Name     string
	Lastname string
}

func NewProfesor(name string, lastname string) Profesor {
	return Profesor{Name: name, Lastname: lastname}
}

func (profesor Profesor) PrintDetails() string {
	return fmt.Sprintf("%v, %v", profesor.Id, profesor.Name)
}
