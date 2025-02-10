package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkNumber(num int) string {
	result := ""
	if num%2 == 0 {
		result = "The given number is even"
	} else {
		result = "The given number is odd"
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give a number: ")

	givenNumberStr, _ := reader.ReadString('\n')
	givenNumberInt, err := strconv.Atoi(strings.Trim(givenNumberStr, "\n"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(checkNumber(givenNumberInt))
}
