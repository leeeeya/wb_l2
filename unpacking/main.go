package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func main() {
	slash := false
	str := `\25\34`
	var res, num string
	var prev, symb rune

	for i, r := range str {
		tmp := string(r)
		fmt.Println(tmp)

		if unicode.IsNumber(r) && i == 0 {
			log.Fatalln("Error: invalid string")
		}

		if (unicode.IsGraphic(r) && !unicode.IsNumber(r)) && (unicode.IsNumber(prev) && !slash) {
			if len(num) == 0 {
				res += string(symb)
			} else {
				if n, err := strconv.Atoi(num); err == nil {
					for j := 0; j < n; j++ {
						res += string(symb)
					}
				} else {
					log.Fatalln("Error: Atoi")
				}
				num = ""
			}
			if r != '\\' {
				symb = r
				prev = r
			} else {
				slash = true
			}
			continue
		}

		if r == '\\' {
			if slash == false {
				if i != 0 {
					res += string(symb)
				}
				slash = true
			} else {
				symb = '\\'
				slash = false
			}
			continue
		}

		if (unicode.IsGraphic(r) && !unicode.IsNumber(r)) && (unicode.IsGraphic(prev) && !unicode.IsNumber(prev)) {
			res += string(symb)
			prev = r
			symb = r
			continue
		}

		if unicode.IsGraphic(r) && !unicode.IsNumber(r) {
			prev = r
			symb = r
			continue
		}

		if unicode.IsNumber(r) && slash {
			prev = r
			symb = r
			slash = false
			continue
		}

		if unicode.IsNumber(r) {
			num += string(r)
			prev = r
		}
	}

	// обработка последнего элемента
	if unicode.IsNumber(prev) && !slash {
		if n, err := strconv.Atoi(num); err == nil {
			for j := 0; j < n; j++ {
				res += string(symb)
			}
		} else {
			log.Fatalln("Error: Atoi")
		}
	} else if unicode.IsNumber(prev) && slash {
		res += string(prev)
	} else if unicode.IsGraphic(prev) {
		res += string(prev)
	}

	fmt.Println(res)
}
