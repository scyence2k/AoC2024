package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type textGrid struct{
	lowerGridBound int
	upperGridBound int
	letterIndexes map[int]string
}

func openFile(s string) *os.File {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	return file
}

func (textGrid *textGrid) checkNonDiagonal(grid [][]string, outterIndex, innerIndex int) int {
	validStrings := 0
	checkPassed := true

	letterIndexLen := len(textGrid.letterIndexes)

	// check left
	for index := 0; index < letterIndexLen; index++ {
		if innerIndex + (letterIndexLen - 1)  > textGrid.upperGridBound {
			checkPassed = false
			break
		}
		if grid[outterIndex][innerIndex + index] != textGrid.letterIndexes[index] {
			checkPassed = false
			break
		}
	}
	if checkPassed {
		validStrings++
	} else {
		checkPassed = true
	}

	//check right
	for index := 0; index < letterIndexLen; index++ {
		if innerIndex - (letterIndexLen - 1) < textGrid.lowerGridBound {
			checkPassed = false
			break
		}
		if grid[outterIndex][innerIndex - index] != textGrid.letterIndexes[index] {
			checkPassed = false
			break
		}
	}
	if checkPassed {
		validStrings++
	} else {
		checkPassed = true
	}

	//check up
	for index := 0; index < letterIndexLen; index++ {
		if outterIndex - (letterIndexLen - 1) < textGrid.lowerGridBound {
			checkPassed = false
			break
		}
		if grid[outterIndex - index][innerIndex] != textGrid.letterIndexes[index] {
			checkPassed = false
			break
		}
	}
	if checkPassed {
		validStrings++
	} else {
		checkPassed = true
	}

	//check down
	for index := 0; index < letterIndexLen; index++ {
		if outterIndex + (letterIndexLen - 1) > textGrid.upperGridBound {
			checkPassed = false
			break
		}
		if grid[outterIndex + index][innerIndex] != textGrid.letterIndexes[index] {
			checkPassed = false
			break
		}
	}
	if checkPassed {
		validStrings++
	} else {
		checkPassed = true
	}

	return validStrings
}

func (textGrid *textGrid) checkDiagonal(grid [][]string, outterIndex, innerIndex int) int {

	validStrings := 0
	checkPassed := true

	letterIndexLen := len(textGrid.letterIndexes)

	// up left
	for index := 0; index < letterIndexLen; index++ {
		if outterIndex - (letterIndexLen - 1) < textGrid.lowerGridBound || innerIndex + (letterIndexLen - 1) > textGrid.upperGridBound {
			checkPassed = false
			break
		}
		if grid[outterIndex - index][innerIndex + index] != textGrid.letterIndexes[index] {
			checkPassed = false
			break
		}
	}
	if checkPassed {
		validStrings++
	} else {
		checkPassed = true
	}

	// up right
	for index := 0; index < letterIndexLen; index++ {
		if outterIndex - (letterIndexLen - 1) < textGrid.lowerGridBound || innerIndex - (letterIndexLen - 1) < textGrid.lowerGridBound {
			checkPassed = false
			break
		}
		if grid[outterIndex - index][innerIndex - index] != textGrid.letterIndexes[index] {
			checkPassed = false
			break
		}
	}
	if checkPassed {
		validStrings++
	} else {
		checkPassed = true
	}

	// down left
	for index := 0; index < letterIndexLen; index++ {
		if outterIndex + (letterIndexLen - 1) > textGrid.upperGridBound || innerIndex + (letterIndexLen - 1) > textGrid.upperGridBound {
			checkPassed = false
			break
		}
		if grid[outterIndex + index][innerIndex + index] != textGrid.letterIndexes[index] {
			checkPassed = false
			break
		}
	}
	if checkPassed {
		validStrings++
	} else {
		checkPassed = true
	}

	// down right
	for index := 0; index < letterIndexLen; index++ {
		if outterIndex + (letterIndexLen - 1) > textGrid.upperGridBound || innerIndex - (letterIndexLen - 1) < textGrid.lowerGridBound {
			checkPassed = false
			break
		}
		if grid[outterIndex + index][innerIndex - index] != textGrid.letterIndexes[index] {
			checkPassed = false
			break
		}
	}
	if checkPassed {
		validStrings++
	} else {
		checkPassed = true
	}

	return validStrings
}

func main() {
	file := openFile("puzzle4.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var xmasText [][]string
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), "")

		xmasText = append(xmasText, temp)
	}

	m := map[int]string {
		0: "X",
		1: "M",
		2: "A",
		3: "S",
	}

	xmasTextGrid := &textGrid{
		lowerGridBound: 0,
		upperGridBound: len(xmasText) - 1,
		letterIndexes: m,
	}

	matches := 0
	for row := 0; row < len(xmasText); row++ {
		for col := 0; col < len(xmasText[row]); col++ {
			if xmasText[row][col] == m[0] {
				matches += xmasTextGrid.checkNonDiagonal(xmasText, row, col)
				matches += xmasTextGrid.checkDiagonal(xmasText, row, col)
			}
		}
	}

	fmt.Println(matches)
}
