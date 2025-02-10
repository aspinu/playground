package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//
// func splitString (givenStrign string) {
//
//   sliceString := strings.Trim (givenStrign, " ")
//
//   return sliceString
// }

func main() {
	fmt.Println("Type a random sentance: ")
	reader := bufio.NewReader(os.Stdin)

	readElement, _ := reader.ReadString('\n')
	formatedElement := strings.Trim(readElement, "\n")
	ready := strings.Split(formatedElement, " ")
	fmt.Println(ready)
}
