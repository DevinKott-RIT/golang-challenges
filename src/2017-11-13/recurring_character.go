package main

import (
	"fmt"
	"errors"
)

func main() {
	inputs := []string {
		"ABCDEBC",
		"IKEUNFUVFV",
		"PXLJOUDJVZGQHLBHGXIW",
		"*l1J?)yn%R[}9~1\"=k7]9;0[$",
	}

	for _, input := range inputs {
		index, err := findFirstRecurringChar(input)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%d (%c)\n", index, input[index])
		}
	}
}

func findFirstRecurringChar(input string) (int, error) {
	array := make([]bool, 128)
	for index, char := range input {
		c := int(char)
		if array[c] == true {
			for i := index - 1; i >= 0; i-- {
				if input[i] == input[index] {
					return i, nil
				}
			}
		} else {
			array[c] = true
		}
	}
	return -1, errors.New("could not find a recurring character")
}