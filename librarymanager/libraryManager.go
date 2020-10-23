package librarymanager

import "fmt"

func LibrayManager(input string) {
	fmt.Println("[Libray Manager]")
	fmt.Println("Input : ", input)
	output, _ := BooksShorter(input)
	fmt.Println("Ouput : ", output)
}
