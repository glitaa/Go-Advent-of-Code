package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var Err2020NotFound = errors.New("couldn't find numbers that summed give 2020")
var ErrNotANumber = errors.New("didn't get a number but wanted one")

type Year2020 struct {
	Multiplication int
	Numbers        []int
}

func (y *Year2020) FindWithTwoNumbers(numbers []int) error {
	for i, num1 := range numbers {
		for j, num2 := range numbers {
			if i == j {
				continue
			}
			if num1+num2 == 2020 {
				*y = Year2020{num1 * num2, []int{num1, num2}}
				return nil
			}
		}
	}
	return Err2020NotFound
}

func (y *Year2020) FindWithThreeNumbers(numbers []int) error {
	for i, num1 := range numbers {
		for j, num2 := range numbers {
			for k, num3 := range numbers {
				if i == j {
					break
				}
				if i == k || j == k {
					continue
				}
				if num1+num2+num3 == 2020 {
					*y = Year2020{num1 * num2 * num3, []int{num1, num2, num3}}
					return nil
				}
			}
		}
	}
	return Err2020NotFound
}

func main() {
	numbers := readNumbers()
	myYear2020 := Year2020{}

	/*myYear2020.FindWithTwoNumbers(numbers)
	if err == nil {
		fmt.Printf("Multiplication of two numbers that summed give 2020: %d \nThe three numbers are: %d, %d, %d\n", myYear2020.Multiplication, myYear2020.Numbers[0], myYear2020.Numbers[1], myYear2020.Numbers[2])
	} else {
		log.Fatalf("%s, given numbers: %v", err, numbers)
	}*/

	findErr := myYear2020.FindWithThreeNumbers(numbers)
	if findErr == nil {
		fmt.Printf("Multiplication of three numbers that summed give 2020: %d \nThe three numbers are: %d, %d, %d\n", myYear2020.Multiplication, myYear2020.Numbers[0], myYear2020.Numbers[1], myYear2020.Numbers[2])
	} else {
		log.Printf("%s, given numbers: %v", findErr, numbers)
	}
}

func readNumbers() (numbers []int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("%s, error: %q", ErrNotANumber, err)
		}
		numbers = append(numbers, number)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return
}