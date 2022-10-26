package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type flags struct {
	fields    string
	delimiter string
	separated bool
}

func (flagStruct *flags) flagParsing() {
	flag.StringVar(&flagStruct.fields, "f", "", "Выбрать поля (колонки)")
	flag.StringVar(&flagStruct.delimiter, "d", "\t", "Выбрать другой разделитель (TAB по умолчанию)")
	flag.BoolVar(&flagStruct.separated, "s", false, "Только строки c разделителем")
	flag.Parse()
}

func main() {
	res := [][]string{}
	flags := &flags{}
	flags.flagParsing()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Ввод:")
		ok := scanner.Scan()
		if !ok && scanner.Err() == nil {
			break
		}

		str := scanner.Text()
		if len(str) == 0 {
			break
		}

		if flags.separated {
			if !strings.Contains(str, flags.delimiter) {
				continue
			}
		}
		res = append(res, strings.Split(str, flags.delimiter))
	}
	if flags.fields != "" {
		nf := strings.Split(flags.fields, ",")
		for _, val := range nf {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Println(err)
			}
			if num-1 >= len(res) {
				continue
			}
			fmt.Println(res[num-1])
		}
		return
	}
	fmt.Println(res)
}
