package polindrom

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func validateInput(start, end float64) (validate bool, err error) {
	minimumValue := float64(1)
	maximumValue := math.Pow(10, 9)
	err = errors.New("The start and the end number is not valid input, check your input value again, please !")
	if (start < end) && (start >= minimumValue && start <= maximumValue) && (end >= minimumValue && end <= maximumValue) {
		validate = true
		err = nil
	}
	return
}

func countPolindromNumbers(start, end float64) (float64, error) {
	validate, err := validateInput(start, end)
	totalPolindNumber := float64(0)
	if validate {
		for number := start; number <= end; number++ {
			strNumber := strconv.FormatFloat(number, 'f', 0, 64)
			riverseStrNumber := reverse(strNumber)
			if strings.Compare(strNumber, riverseStrNumber) == 0 {
				totalPolindNumber++
			}
		}
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
	if len(input) > 0 {
		splitedString := strings.Split(input, " ")
		if len(splitedString) != 2 {
			return 0, errors.New("Input value must a given string with 2 numbers that split by space ( )")
		}
		startNumb, _ := strconv.ParseFloat(splitedString[0], 6)
		endNumb, _ := strconv.ParseFloat(splitedString[1], 6)
		fmt.Println("Input : ", input)
		output, err := countPolindromNumbers(startNumb, endNumb)
		fmt.Println("Output : ", output)
		return output, err
	}
	return 0, errors.New("Input must be a string, not null or an empty string")

}
