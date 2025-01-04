package main

import "fmt"

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("Cannot divide %v by 0 . \n", de.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func FindQ(dividend, divisor float64) {
	q, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The qoutient is %.2f .\n", q)
}

func ErrorTest() {
	FindQ(10, 5)
	FindQ(10, 0)
}
