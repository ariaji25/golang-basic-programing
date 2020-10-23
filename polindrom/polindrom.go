package polindrom

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

/* Reverse String Function.
Function to reverse a string.
- Input : A Sting
- Output : A reverse of input string
*/
func reverse(str string) (result string) {
	// do looping to get each characters in string
	for _, v := range str {
		// add the caracter to new result string by added it as the first character, so the result is reverse of the input string
		result = string(v) + result
	}
	return
}

/* Validate Input Function.
Function to validate the input for this program,
the input will valid if the start number is less than end number
and the manimum value of the start and end number is 1
and the maximum value of the start and the end number is 10^9.
*/
// -	Input :
// 	- start : Is the start number of range numbers.
// 	- end 	: Is the end  number of range numbers.
// -	Output :
// 	- validate : witll true if input is valid, and false if not valid
// 	- err : return and error if the inpus is not valid
func validateInput(start, end float64) (validate bool, err error) {
	// init the minimum value of input
	minimumValue := float64(1)
	// init the maximum value of input
	maximumValue := math.Pow(10, 9)
	// init error message if input is not valid
	err = errors.New("The start and the end number is not valid input, check your input value again, please !")
	// chekc if input is valid
	if (start < end) && (start >= minimumValue && start <= maximumValue) && (end >= minimumValue && end <= maximumValue) {
		// set input is valid
		validate = true
		// set error is null
		err = nil
	}
	return
}

/*
Count Polindrom Function.
A function to count how many polindrom numbers in the range between two numbers :
*/
// -	Input :
// 	- start : Is the start number of range numbers.
// 	- end 	: Is the end  number of range numbers.
// 	- example : 1,10 99,100
// - Output :
// 	- total of polindrom numbers
// 	- return an error if there is an erro while find the polindrom number
func countPolindromNumbers(start, end float64) (float64, error) {
	// validate the input
	validate, err := validateInput(start, end)
	// init the total of polindrom numbers
	totalPolindNumber := float64(0)
	// check if input is valid
	if validate {
		// count the polindrom number by looping until the end range of numbers.
		for number := start; number <= end; number++ {
			// convert number value to string value
			strNumber := strconv.FormatFloat(number, 'f', 0, 64)
			// get reverse string of converted number above
			riverseStrNumber := reverse(strNumber)
			// chekc the number is polindrom if the converted number above is equal to it reverse string
			if strings.Compare(strNumber, riverseStrNumber) == 0 {
				// icrese total of polindrom numbers
				totalPolindNumber++
			}
		}
		// set error is null to set the program is not error
		err = nil
	}
	return totalPolindNumber, err
}

/*
PolindromNumber function.
A Function to count how many polindrom numbers between two range of numbers from a given input string.
*/
// -	input 	:	The input must be a String, and the value must be two given number that split by space,
// 					start number must be greater and less than end number, the minimum value of number is one
//  				and the maximum value is less than or equals to 10^9.
// 					example : "1 10", "99 100", "21 33".
// -	output 	: 	The output is a number with type float64, which is the total of polindrom numbers, and if there is
// 			  		an error it will given an error with message.
func PolindromNumber(input string) (float64, error) {
	fmt.Println("[Polindrom Numbers]")
	// chekc if input is the correct input, not and empty string
	if len(input) > 0 {
		// parse the input string to list, split by space
		splitedString := strings.Split(input, " ")
		// check the size of list, if size is not 2, the input is incorrect
		if len(splitedString) != 2 {
			return 0, errors.New("Input value must a given string with 2 numbers that split by space ( )")
		}
		// get the start and the end number to find the range of numbers
		startNumb, _ := strconv.ParseFloat(splitedString[0], 6)
		endNumb, _ := strconv.ParseFloat(splitedString[1], 6)
		// show the input value
		fmt.Println("Input : ", input)
		// get total of polindrom Numbers
		output, err := countPolindromNumbers(startNumb, endNumb)
		// show the output value
		fmt.Println("Output : ", output)
		return output, err
	}
	return 0, errors.New("Input must be a string, not null or an empty string")

}
