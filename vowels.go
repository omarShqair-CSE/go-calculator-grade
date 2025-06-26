// Write a function CountVowels(s string) int that returns the number of vowels (a, e, i, o, u) in the string s.
// Test it with "Hello World" and "Go is great!".

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CountVowels(s string) int {
	s = strings.ToLower(s)
	count := 0
	//a, e, i, o, u
	for _, data := range s {
		switch data {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your sentence: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	fmt.Println("Number of vowels:", CountVowels(userInput))
}
