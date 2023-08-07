package main

import (
	"fmt"
	"go/src/keyboard"
	"log"
)

const (
	INVALID_INPUT = "Invalid input"
)

func main() {
	checkGrade()
	convCelToFht()
}

func checkGrade() {
	fmt.Printf("Enter a grade : ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(INVALID_INPUT, err)
	}

	var status string
	if grade >= 60 {
		status = "passed"
	} else {
		status = "failed"
	}

	fmt.Println("A grade of ", grade, "is ", status)
}

func convCelToFht() {
	fmt.Print("Enter temp in Celcius : ")
	celcius, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(INVALID_INPUT, err)
	}
	fehrenheit := (celcius + 32) * 9 / 5
	fmt.Printf("%0.2f celcius -> %0.2f fehrenheit", celcius, fehrenheit)
}
