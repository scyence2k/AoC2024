package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func str_to_int(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err.Error())
	}

	return num
}

func Run() {
	file, err := os.Open("puzzle1.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)

	var lList []int
	var rList []int
	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Split(line, "   ")
		lList = append(lList, str_to_int(values[0]))
		rList = append(rList, str_to_int(values[1]))
	}

	sort.Ints(lList)
	sort.Ints(rList)

	total_distance := 0
	for i, num := range lList {
		diff := rList[i] - num
		if diff < 0 {
			diff = diff * -1
		}

		total_distance += diff
	}
	fmt.Printf("%d\n", total_distance)

	similarity_score := 0
	for _, lListElem := range lList {
		match_amount := 0
		for _, rListElem := range rList {
			if lListElem == rListElem {
				match_amount++
			}
		}
		similarity_score += lListElem * match_amount
	}
	fmt.Printf("%d\n", similarity_score)
}
