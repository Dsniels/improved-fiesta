package entities

import "fmt"

type Curso struct {
	Id          int
	Name        string
	Description string
	Profesor    Profesor
}

func (c Curso) String() string {
	return fmt.Sprintf("Id: %v, Curso: %v, Description: %v - Profesor: %v",c.Id, c.Name, c.Description, c.Profesor)
}