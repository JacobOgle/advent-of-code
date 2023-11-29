package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Takes the day of the input and grabs the contents -> stored in byte array
func ByteArrayParseInput(day int) []byte {
	dayToString := fmt.Sprint(day)
	data, err := os.ReadFile("../inputs/d" + dayToString + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// Uses bufio to parse a file line by line
// Assumes the input is integers w/o delimiters
// procFunc is a processing function for the data
// Day 1 use, but not sure if I can generalize this
func BufioParseIntegerInput(day int, procFunc func()) {
	readFile, err := os.Open("../inputs/d" + fmt.Sprint(day) + ".txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			continue
		}

		n, e := strconv.Atoi(fileScanner.Text())
		if e != nil {
			log.Fatal(e)
		}

		fmt.Println(n)
	}

	readFile.Close()
}
