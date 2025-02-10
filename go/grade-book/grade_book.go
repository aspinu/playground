package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func avrg(list []int) string {
	totlaScore := 0
	for _, number := range list {
		totlaScore += number
	}

	avrageScore := totlaScore / int(len(list))
	grade := ""
	if avrageScore <= 100 && avrageScore >= 90 {
		grade = "A"
	} else if avrageScore < 90 && avrageScore >= 80 {
		grade = "B"
	} else if avrageScore < 80 && avrageScore >= 70 {
		grade = "C"
	} else {
		grade = "D"
	}

	return grade
}

func main() {
	pr := fmt.Println
	reader := bufio.NewReader(os.Stdin)
	pr("Give scores: \n ")
	scores := []int{}
	for i := 0; i < 3; i++ {
		str_element, _ := reader.ReadString('\n')
		int_element, err := strconv.Atoi(strings.Trim(str_element, "\n"))
		if err == nil {
			scores = append(scores, int_element)
		} else {
			log.Fatal(err)
		}

	}
	yourGreade := avrg(scores)
	pr("Your Score is: ", yourGreade)
}
