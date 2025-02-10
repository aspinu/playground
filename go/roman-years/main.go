package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	rom "github.com/brandenc40/romannumeral"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give a year to convert to Roman notation: ")

	yearStr, _ := reader.ReadString('\n')

	yearInt, err := strconv.Atoi(strings.Trim(yearStr, "\n"))
	if err != nil {
		log.Fatal(err)
	}

	romanic, err := rom.IntToString(yearInt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The year in Romna numberal is: %s\n", romanic)
}
