package main

import "fmt"

func main() {
	sentences := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	freq := func(char string) map[string]int {
		chars := make(map[string]int)
		for _, char := range char {
			if string(char) != " " && string(char) != "," {
				chars[string(char)]++
			}
		}
		return chars
	}

	func(sentences string) {
		sum := make(map[string]int)
		res := make(chan map[string]int, len(sentences))
		for _, char := range sentences {
			go func(char string) {
				res <-freq(char)
			}(string(char))
		}

		for i := 0; i < len(sentences); i++ {
			for char, n := range <-res {
				sum[char] += n
				fmt.Printf("%s: %v\n", char, sum[char])
			}
		}
	}(sentences)
}
