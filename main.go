package main

import (
	"fmt"
	"os"
	"github.com/Dsniels/improved-fiesta/utils"
)

func main() {
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
