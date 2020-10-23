package lostnumber

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func findlostNumber(numbers string, startnumber int) (lost float64, err error) {
	minimumValue := float64(1)
	maximumValue := math.Pow(10, 6)
	strnumber := strconv.Itoa(startnumber)
	for i := float64(startnumber) + 1; i > minimumValue && i <= maximumValue; i++ {
		strnumber = strnumber + strconv.FormatFloat(i, 'f', 0, 64)
		if strings.Compare(numbers, strnumber) == 1 {
			if !strings.Contains(numbers, strnumber) {
				if lost == 0 {
					lost = i
					err = nil
				} else {
					i = maximumValue
				}
			}
		} else {
			lost = 0
			i = maximumValue
			err = errors.New("Expected sorted list of numbers with a lost number, Get a sorted list without lost number")
		}
	}
	fmt.Println("Lost Number : ", lost)
	return
}

func findStartNumber(numbers string) (result int, err error) {
	fmt.Println("Input : ", numbers)
	err = errors.New("Input Is Invalid, input must be a list of minimum 3 numbers")
	maxdigit := 7
	for i := maxdigit; i >= 1; i-- {
		if len(numbers) > i {
			startnumber, _ := strconv.Atoi(numbers[:i])
			strFisrtSecondNumber := strconv.Itoa(startnumber) + strconv.Itoa(startnumber+1)
			strFisrtThirdNumber := strconv.Itoa(startnumber) + strconv.Itoa(startnumber+2)
			if strings.Contains(numbers, strFisrtSecondNumber) {
				strFirstForthNumber := strFisrtSecondNumber + strconv.Itoa(startnumber+2)
				strFirstFifthNumber := strFisrtSecondNumber + strconv.Itoa(startnumber+3)
				if strings.Contains(numbers, strFirstForthNumber) || strings.Contains(numbers, strFirstFifthNumber) {
					fmt.Println("Start Number : ", startnumber)
					result = startnumber
					i = 0
					err = nil
				}
			} else if strings.Contains(numbers, strFisrtThirdNumber) {
				fmt.Println("Start Number : ", startnumber)
				result = startnumber
				i = 0
				err = nil
			}
		}
	}
	return
}

func LostNumber(numbers string) (output float64, err error) {
	fmt.Println("Input											| Output")
	startnumber, err := findStartNumber(numbers)
	output, err = findlostNumber(numbers, startnumber)
	fmt.Println("", numbers, " | ", output)
	return
	// exampe input
	// fmt.Println("[Lost Number]")
	// findStartNumber("23242526272830")
	// fmt.Println()
	// findStartNumber("101102103104105106107108109111112113")
	// fmt.Println()
	// findStartNumber("12346789")
	// fmt.Println()
	// findStartNumber("79101112")
	// fmt.Println()
	// findStartNumber("7891012")
	// fmt.Println()
	// findStartNumber("9799100101102")
	// fmt.Println()
	// findStartNumber("100001100002100003100004100006")
}
