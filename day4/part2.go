package day4

import (
	"bufio"
	"fmt"
	_"log"
	_"os"
	"strings"
)


func RunPart2() {
	file := openFile("puzzle4.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var xmasText [][]string
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), "")

		xmasText = append(xmasText, temp)
	}

	m := map[int]string{
		0: "M",
		1: "A",
		2: "S",
	}

	xmasTextGrid := &textGrid{
		lowerGridBound: 0,
		upperGridBound: len(xmasText) - 1,
		letterIndexes:  m,
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
