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
