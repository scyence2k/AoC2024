package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type LineDirectionType int

const (
	Unclear LineDirectionType = iota
	Ascending
	Descending
)

func openFile(s string) *os.File {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	return file
}

func getLevelDifference(num1, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		diff *= -1
	}

	return diff
}

func reportIsSafe(report []int) bool {
	DirectionType := Unclear
	prevNumber := 0

	for index, level := range report {
		if index == 0 {
			prevNumber = level
			continue
		}

		if prevNumber == level {
			return false
		}

		diff := getLevelDifference(prevNumber, level)
		if diff > 3 {
			return false
		}

		switch DirectionType {
		case Unclear:
			if prevNumber < level {
				DirectionType = Ascending
			} else {
				DirectionType = Descending
			}
		case Ascending:
			if prevNumber > level {
				return false
			}
		case Descending:
			if prevNumber < level {
				return false
			}
		}

		prevNumber = level
	}

	return true
}

func reportHasSafeVariation(report []int) bool {

	safe := false
	for i := 0; i < len(report); i++ {
		fmt.Printf("%d\n", report)
		reportCopy := slices.Clone(report)
		reportCopy = slices.Delete(reportCopy, i , i+1)
		fmt.Printf("%d\n", reportCopy)

		if reportIsSafe(reportCopy) {
			safe = true
			break
		}
	}

	return safe
}

func Run() {
	file := openFile("puzzle2.txt")
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
		safe := reportIsSafe(report) || reportHasSafeVariation(report)

		if safe {
			safe_reports++
		}

	}

	fmt.Print(safe_reports)

}
