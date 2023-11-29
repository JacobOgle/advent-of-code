package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func DayOnePartOneSoln() {
	res := process(1)
	fmt.Println("D1P1 -- ", res)
}

func check_max(currMax, currSum int) bool {
	return currSum > currMax
}

func process(day int) int {
	var currMax int = 0
	var currSum int = 0
	readFile, err := os.Open("../inputs/d" + fmt.Sprint(day) + ".txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		if fileScanner.Text() == "" {
			swap := check_max(currMax, currSum)
			if swap {
				currMax = currSum
			}
			currSum = 0
			continue
		}

		n, e := strconv.Atoi(fileScanner.Text())
		if e != nil {
			log.Fatal(e)
		}
		currSum += n
	}

	readFile.Close()
	return currMax
}
