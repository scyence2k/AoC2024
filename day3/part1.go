package day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func openFile(s string) *os.File {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	return file
}

func Run() {
	file, err := os.ReadFile("puzzle3.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	validMulStrings := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))`)

	allMulMatches := validMulStrings.FindAllString(string(file), -1)

	total := 0
	for _, match := range allMulMatches {
		numbers := strings.Split(match[4:len(match)-1], ",")
		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		total += num1 * num2
	}

	fmt.Print(total)
}