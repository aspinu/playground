package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	arry := []int{}
	print("Give a number: \n")
	for i := 0; i < 5; i++ {

		str_element, _ := reader.ReadString('\n')
		int_element, err := strconv.Atoi(strings.Trim(str_element, "\n"))
		if err == nil {
			arry = append(arry, int_element)
		} else {
			log.Fatal(err)
		}

	}
	sort.Ints(arry)
	fmt.Println(arry)
	fmt.Println("Smallest number in array is: ", arry[0])
}
