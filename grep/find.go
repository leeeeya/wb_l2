package grepSupport

import (
	"fmt"
	"regexp"
	"strconv"
)

func (flagStruct Flags) FindWithoutFlags(str []string) []string {
	res := make([]string, 0)
	for _, s := range str {
		matched, _ := regexp.MatchString(flagStruct.DesiredValue, s)
		if matched {
			res = append(res, s)
		}
	}
	return res
}

func (flagStruct Flags) FindInvert(str []string) []string {
	res := make([]string, 0)
	for _, s := range str {
		matched, _ := regexp.MatchString(flagStruct.DesiredValue, s)
		if matched {
			continue
		}
		res = append(res, s)
	}
	return res
}

func (flagStruct Flags) FindLineNum(str []string) []string {
	res := make([]string, 0)
	counter := 1
	for _, s := range str {
		matched, _ := regexp.MatchString(flagStruct.DesiredValue, s)
		if matched {
			a := strconv.Itoa(counter) + ": " + s
			res = append(res, a)
			counter++
		}
	}
	return res
}

func (flagStruct Flags) FindFixed(str []string) []string {
	res := make([]string, 0)

	for _, s := range str {
		matched, _ := regexp.MatchString("^"+flagStruct.DesiredValue+"$", s)
		if matched {
			res = append(res, s)
		}
	}

	return res
}

func (flagStruct Flags) FindContStrings(str []string) []string {
	res := make([]string, 0)
	counter := 0
	for _, s := range str {
		matched, _ := regexp.MatchString(flagStruct.DesiredValue, s)
		if matched {
			counter++
		}
	}
	res = append(res, strconv.Itoa(counter))
	return res
}

func (flagStruct Flags) FindAfter(str []string) []string {
	res := make([]string, 0)

	for i := 0; i < len(str); i++ {
		matched, _ := regexp.MatchString(flagStruct.DesiredValue, str[i])
		if matched {
			for j := 0; j < flagStruct.a+1; i, j = i+1, j+1 {
				if i != len(str)-1 {
					res = append(res, str[i])
				} else {
					return res
				}
			}
			i--
		}
	}
	return res
}

func (flagStruct Flags) FindBefore(str []string) []string {
	res := make([]string, 0)
CYCLE:
	for i := len(str) - 1; i > 0; i-- {
		matched, _ := regexp.MatchString(flagStruct.DesiredValue, str[i])
		if matched {
			for ; i > i-flagStruct.b-1; i-- {
				if i != 0 {
					res = append(res, str[i])
				} else {
					break CYCLE
				}
			}
			i++
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func (flagStruct Flags) FindContext(str []string) []string {
	res := make([]string, 0)
	tmpMap := make(map[string]int)
	/*
		Похожу по строке
		Если встречаю нужную, то кладу в res
		каждую добавленную строку кладу в мапу, как значение кладу индекс
	*/
	for i := 0; i < len(str); i++ {
		matched, _ := regexp.MatchString(flagStruct.DesiredValue, str[i])
		if matched {
			goalMin := 0
			goalMax := i + flagStruct.cBig
			if i >= flagStruct.cBig {
				goalMin = i - flagStruct.cBig
			}
			for i = goalMin; i < goalMax+1; i++ {
				if i != len(str) {
					if _, ok := tmpMap[str[i]]; ok && tmpMap[str[i]] == i {
						continue
					} else {
						res = append(res, str[i])
					}
				}
			}
		}
	}
	return res
}

func (flagStruct Flags) Find(str []string) []string {
	flagsOff := true
	res := make([]string, 10)
	if flagStruct.v {
		flagsOff = false
		res = flagStruct.FindInvert(str)
	}

	if flagStruct.n {
		flagsOff = false
		res = flagStruct.FindLineNum(str)
	}

	if flagStruct.f {
		flagsOff = false
		res = flagStruct.FindFixed(str)
	}

	if flagStruct.cSmall {
		flagsOff = false
		res = flagStruct.FindContStrings(str)
	}

	if flagStruct.a > -1 {
		flagsOff = false
		res = flagStruct.FindAfter(str)
	}

	if flagStruct.b > -1 {
		flagsOff = false
		res = flagStruct.FindBefore(str)
	}

	if flagStruct.cBig > -1 {
		flagsOff = false
		res = flagStruct.FindContext(str)
	}

	if flagsOff {
		fmt.Println("?")
	}
	return res
}
