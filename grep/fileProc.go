package grepSupport

import (
	"bufio"
	"log"
	"os"
)

func (flagStruct Flags) OpenFile() []string {
	file, err := os.Open(flagStruct.FileName)
	if err != nil {
		log.Fatal("File not open: ", err)
	}
	fileScanner := bufio.NewScanner(file)
	res := make([]string, 1)
	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		res = append(res, fileScanner.Text())
	}
	return res
}
