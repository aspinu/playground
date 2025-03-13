package main

import "fmt"

func isValid(s string) bool {
	var resSq bool
	var resCi bool
	var resAc bool
	var listC []string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			listC = append(listC, string(s[i]))
			for j := 0; j < len(s); j++ {
				if string(s[j]) == ")" {
					listC = append(listC, string(s[j]))
				}
			}
		}
		if len(listC)%2 == 0 {
			resCi = true
		} else {
			resCi = false
		}

		if string(s[i]) == "[" {
			listC = append(listC, string(s[i]))
			for j := 0; j < len(s); j++ {
				if string(s[j]) == "]" {
					listC = append(listC, string(s[j]))
				}
			}
		}
		if len(listC)%2 == 0 {
			resSq = true
		} else {
			resSq = false
		}
		if string(s[i]) == "{" {
			listC = append(listC, string(s[i]))
			for j := 0; j < len(s); j++ {
				if string(s[j]) == "}" {
					listC = append(listC, string(s[j]))
				}
			}
		}
		if len(listC)%2 == 0 {
			resAc = true
		} else {
			resAc = false
		}

	}
	if resCi && resSq && resAc {
		return true
	} else {
		return false
	}
}

func main() {
	a := "([]){"
	b := isValid(a)
	fmt.Println(b)
}
