package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/spatialcurrent/go-math/pkg/math"
)

func DayOneSoln() {
	var t3sum int = 0
	max, t3 := process(1, t3sum)

	fmt.Println("D1P1 -- ", max)
	fmt.Println("D1P2 -- ", t3)
}

func check_max(currMax, currSum int) bool {
	return currSum > currMax
}

func process(day, t3arr int) (int, int) {
	var currMax int = 0
	var currSum int = 0

	var sumArr []int

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
			sumArr = append(sumArr, currSum)
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

	slices.Sort(sumArr)          // sort array in ascending order
	l3 := sumArr[len(sumArr)-3:] // grab the last 3 elements of the array

	top3, err := math.Sum(l3) // sum the 3 array elements
	if err != nil {
		log.Fatal(err)
	}
	return currMax, top3.(int) // for top3 we need to cast to int in a interface specific manner
}
