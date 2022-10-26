package main

import (
	grepSupport "burmachine/grep-support/grep-support"
	"fmt"
	"os"
)

func main() {
	Flags := &grepSupport.Flags{FileName: os.Args[len(os.Args)-1]}
	Flags.Input()
	str := Flags.OpenFile()
	str2 := Flags.Find(str)
	for i, s := range str2 {
		if i == 0 && s == "" {
			continue
		}
		fmt.Println(s)
	}
	fmt.Println(Flags)
}
