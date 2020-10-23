package lostnumber

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

/* Find Lost Number.
Function to fun lost number from a sorted list of numbers by given the list numbers and the first number in list.
- input :
	- numbers : String that contains the sorted list numbers.
	- startnumber : is the first number in the list numbers.
- Output :
	- lost : is the lost number in the list numbers.
	- err : an error if there is an error while find the lost number.*/
func findlostNumber(numbers string, startnumber int) (lost float64, err error) {
	// init the minimum value of number
	minimumValue := float64(1)
	// init the maximum value of number
	maximumValue := math.Pow(10, 6)
	// convert the first number to string
	strnumber := strconv.Itoa(startnumber)
	// do looping to find the lost number, by increment the first number and find it in the list numbers
	for i := float64(startnumber) + 1; i > minimumValue && i <= maximumValue; i++ {
		// join the string of first number with the next value from incremente the first number
		strnumber = strnumber + strconv.FormatFloat(i, 'f', 0, 64)
		// find if the joined string of number above is not equal to string that contains list of numbers
		if strings.Compare(numbers, strnumber) == 1 {
			// check if the joined number before, is it contains in the string that contains list of numbers
			if !strings.Contains(numbers, strnumber) {
				// check if lost number is equal to 0
				if lost == 0 {
					// set the lost number
					lost = i
					// set error is null
					err = nil
				} else {
					// stop the loop by set "i" is reach the maximum value
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

/* FIndStartNumber Function.
Function to find the start number from the sorted list number in string.
-	input : string that contains the sorted list of numbers
- 	output : the result number with type integer and the error if there is and error while find the start number*/
func findStartNumber(numbers string) (result int, err error) {
	// init the error if filed to find the start number
	err = errors.New("Input Is Invalid, input must be a list of minimum 3 numbers")
	// init maximum digit for numbers
	maxdigit := 7
	// do looping with the maximum digit as the range
	for i := maxdigit; i >= 1; i-- {
		// check if len of list numbers is more than the digits of number in "i"
		if len(numbers) > i {
			// get the first number that converted from the first number with type string from list
			firstnumber, _ := strconv.Atoi(numbers[:i])
			// define the secon number, convert it to string and join it with the first number
			strFisrtSecondNumber := strconv.Itoa(firstnumber) + strconv.Itoa(firstnumber+1)
			// define the thir number, convert it to string and join it with the start number
			strFisrtThirdNumber := strconv.Itoa(firstnumber) + strconv.Itoa(firstnumber+2)
			// check if joined first and second number is contain in the string that contain the list numbers
			if strings.Contains(numbers, strFisrtSecondNumber) {
				// get the fourth number of list numbers, convert it to string and join it with the first number, second number
				strFirstForthNumber := strFisrtSecondNumber + strconv.Itoa(firstnumber+2)
				// get the fifth number of list numbers, convert it to string and join it with the first number, second number
				strFirstFifthNumber := strFisrtSecondNumber + strconv.Itoa(firstnumber+3)
				// chekc if joined first, second, and four ro fine number is contains in list numbers
				if strings.Contains(numbers, strFirstForthNumber) || strings.Contains(numbers, strFirstFifthNumber) {
					// set the firstnumber
					result = firstnumber
					// stop the loopby by set i equal 0
					i = 0
					// set error is null
					err = nil
				}
				// check if joined first and third number is contains in list numbers
			} else if strings.Contains(numbers, strFisrtThirdNumber) {
				// set the first number
				result = firstnumber
				// stop loop by set i equal to 0
				i = 0
				// set error is null
				err = nil
			}
		}
	}
	return
}

/*
Lost Number Function.
A Program to find a Lost Number in the list of sorted numbers that contains in given string. Minimum numbers in list is 3 numbers, and the maximum numbers is 10^6 numbers.
The lost numbers is in the center of list not in the start or the end of list.
*/
// -	input : A string that contains list of sorted numbers, examples "124", "121315".
// -	output : the lost number in the list numbers.
// -	err : the error if there is an error while find the lost number.
func LostNumber(numbers string) (output float64, err error) {
	// show the input value
	fmt.Println("Input : ", numbers)
	// find the start number of list number from input value
	startnumber, err := findStartNumber(numbers)
	// get the lost number from the list by given the startnumber of list
	output, err = findlostNumber(numbers, startnumber)
	// show the output
	fmt.Println("OUtput : ", output)
	return
}
