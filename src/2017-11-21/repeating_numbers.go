package main

import "fmt"

func main() {
	inputs := []string{
		"82156821568221",
		"11111011110111011",
		"98778912332145",
		"124489903108444899",
	}

	for _, input := range inputs {
		m := make(map[int]map[string]int)
		strLen := len(input)
		for substringLength := 2; substringLength < strLen; substringLength++ {
			for pos := 0; pos < strLen-substringLength+1; pos++ {
				chunk := input[pos:pos+substringLength]
				if m[substringLength] == nil {
					m[substringLength] = make(map[string]int)
					m[substringLength][chunk] = 1
				} else {
					m[substringLength][chunk]++
				}
			}
		}

		fmt.Printf("%s\n\t", input)
		for i := strLen - 1; i >= 2; i-- {
			if m[i] == nil {
				continue
			}
			for chunk, times := range m[i] {
				if times == 1 {
					continue
				}
				fmt.Printf("%s:%d ", chunk, times)
			}
		}
		fmt.Printf("\n\n")
	}
}
