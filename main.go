package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	currentTime := time.Now()
	currentYear := currentTime.Year()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your current age? ")
	currentAge, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't get current age.")
		return
	}
	currentAge = strings.TrimSpace(currentAge)
	currentAgeInt, err := strconv.Atoi(currentAge)
	if err != nil {
		fmt.Print("Couldn't convert Age to Int")
		return
	}

	fmt.Print("What Age would you like to retire? ")
	retirementAge, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't get retirement age.")
		return
	}
	retirementAge = strings.TrimSpace(retirementAge)
	retirementAgeInt, err := strconv.Atoi(retirementAge)
	if err != nil {
		fmt.Print("Couldn't convert retirement")
		return
	}

	yearsToRetire := retirementAgeInt - currentAgeInt
	yearOfRetirement := currentYear + yearsToRetire

	fmt.Printf("You have %v years left until you can retire\n", yearsToRetire)
	fmt.Printf("Its %v, so you can retire in %v\n", currentYear, yearOfRetirement)

}
