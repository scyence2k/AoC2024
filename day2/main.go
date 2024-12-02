package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkAsc(num1, num2 int) bool {
	return num1 != num2 && num1 < num2 && num2 - num1 > 3
} 

func checkDesc(num1, num2 int) bool {
	return num1 != num2 && num1 > num2 && num1 - num2 > 3
}

func main() {
	file, err := os.Open("./day2/puzzle.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports [][]int
	for scanner.Scan() {
		reportLine := scanner.Text()
		levels := strings.Split(reportLine, " ")

		var levelsInt []int
		for _, level := range levels {
			temp, err := strconv.Atoi(level)
			if err != nil {
				log.Fatal(err.Error())
			}

			levelsInt = append(levelsInt, temp)
		}

		reports = append(reports, levelsInt)
	}

	safe_reports := 0
	for _, report := range reports {
		skip_used := false

		asc_levels := false
		for i := 0; i < len(report) - 1; i++ {
			result := checkAsc(report[i], report[i+1])
			if !result && skip_used {
				asc_levels = false
				break
			}
			if !result && i == 0 && checkAsc(report[i], report[i+2]) {
				asc_levels = true
				skip_used = true
				continue
			}
			if !result && i > 0 && checkAsc(report[i-1], report[i+1]) {
				asc_levels = true
				skip_used = true
				continue
			}

			asc_levels = true
		}

		skip_used = false
		desc_levels := false
		for i := 0; i < len(report) - 1; i++ {
			result := checkDesc(report[i], report[i+1])
			if !result && skip_used {
				asc_levels = false
				break
			}
			if !result && i == 0 && checkDesc(report[i], report[i+2]) {
				asc_levels = true
				skip_used = true
				continue
			}
			if !result && i > 0 && checkDesc(report[i-1], report[i+1]) {
				asc_levels = true
				skip_used = true
				continue
			}
			desc_levels = true
		}

		fmt.Printf("%d, asc: %t, desc: %t\n", report, asc_levels, desc_levels)

		if asc_levels || desc_levels {
			safe_reports++
		}
		
	}

	fmt.Print(safe_reports)

}