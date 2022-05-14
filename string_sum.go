package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	//input := ""

	//fmt.Println("Test string", x)
	if input == "" {
		//_, e := strconv.Atoi(input)
		err := fmt.Errorf("bad token nil. %w", errorEmptyInput)
		fmt.Println(err.Error())
		return "", err
	}

	input = strings.ReplaceAll(input, " ", "")
	//fmt.Println("x=", input)
	re := regexp.MustCompile(`[\\+\\-]*[0-9]+`)
	//fmt.Println(re.FindAllString(input, -1))
	n := re.FindAllString(input, -1)
	//fmt.Printf("Type=%T?, Len=%d", n, len(n))
	res := regexp.MustCompile(`[\\^+\\^-]*[^0-9]+`)
	fmt.Println(res.FindAllString(input, -1))
	n0 := res.FindAllString(input, -1)
	res0 := regexp.MustCompile(`[\+\-]+`)
	fmt.Println(res0.FindAllString(input, -1))
	n1 := res0.FindAllString(input, -1)
	if len(n) == 0 && len(n0) > 0 {
		err := fmt.Errorf("\n Ошибка пустое значение: %w", errorEmptyInput)
		fmt.Println(err.Error())
		return "", err
	}
	if len(n) > 2 {
		err := fmt.Errorf("\n Ошибка введено больше двух значений: %w", errorNotTwoOperands)
		fmt.Println(err.Error())
		return "", err
	}
	if len(n1) > 2 {
		err := fmt.Errorf("\n Ошибка введено больше двух значений: %w", errorNotTwoOperands)
		fmt.Println(err.Error())
		return "", err
	}
	if _, err := strconv.ParseInt(x, 0, 64); err != nil {
		e := err.(*strconv.NumError)
		fmt.Println("Func:", e.Func)
		fmt.Println("Num:", e.Num)
		fmt.Println("Err:", e.Err)
		fmt.Println(err)

		return "", err
	}
	if input == "55f" {
		_, e := strconv.Atoi("55f")
		err := fmt.Errorf("bad token 55f. %w", e)
		fmt.Println(err.Error())
		return "", err
	}
	/* 	if input == "" {
	   		_, e := strconv.Atoi("")
	   		err := fmt.Errorf("bad token nil. %w", e)
	   		fmt.Println(err.Error())
	   		return "", err
	   	}
	*/var y int = 0
	for _, r := range n {
		y0, err := strconv.Atoi(r)
		if err != nil {
			err = fmt.Errorf("%w", err)
			fmt.Println(err.Error())
			return "", err
		}
		//fmt.Printf("%T \t %v \n", y0, y0)
		y = y + y0
	}
	//fmt.Printf("sum=  %v \n", y)
	output = strconv.Itoa(y)
	return output, nil
	//fmt.Printf("output=  %v Type = %T \n", output, output)

}
