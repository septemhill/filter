package main

import (
	"fmt"

	"br"
	"enum"
)

func main() {
	// Case 1: Both of cases without return value, just action
	br.Branch(
		func() (enum.SumType, bool, error) {
			return nil, true, nil
		},
		func(_ enum.SumTypeZero) enum.SumTypeZero {
			fmt.Println("This is if statement")
			return enum.NewSumTypeZero()
		},
		func(_ enum.SumTypeZero) enum.SumTypeZero {
			fmt.Println("This is else statement")
			return enum.NewSumTypeZero()
		},
	)

	// Case 2:
	value := 101
	br.Branch(
		func() (enum.SumType, bool, error) {
			if value%2 == 1 {
				return enum.NewSumTypeOne(value), true, nil
			}
			return enum.NewSumTypeOne(value), false, nil
		},
		func(st enum.SumTypeOne[int]) enum.SumTypeOne[int] {
			return enum.NewSumTypeOne(123)
		},
		func(st enum.SumTypeOne[int]) enum.SumTypeOne[string] {
			return enum.NewSumTypeOne("Septem")
		},
	)
}
