package main

import (
	"fmt"
	"github.com/Dsniels/improved-fiesta/entities"
	"github.com/Dsniels/improved-fiesta/utils"
	"os"
)

func playWithFiles() {

	dir, _ := os.Getwd()
	filepath := dir + "/data/file.txt"

	content, err := utils.ReadTextFile(filepath)

	if err != nil {
		panic("AHHHHHH")
	} else {
		fmt.Printf("Value: %v", content)
		ctn := fmt.Sprintf("File %v\n says %v", filepath, content)
		utils.WriteFile(filepath+".output.txt", ctn)
	}

}

func main() {
	user := entities.NewProfesor("Daniel", "Salazar")
	curso := entities.Curso{Name: "Go", Description: "Aprende go", Profesor: user}
	fmt.Printf("%v", curso.String())

}
