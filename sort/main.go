package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Info struct {
	k        int
	n        bool
	r        bool
	u        bool
	fileName string
}

func (flags Info) openFile() *os.File {
	file, err := os.Open(flags.fileName)
	if err != nil {
		log.Fatal("File not open: ", err)
	}
	return file
}

func (flags *Info) fileProc(file *os.File) []string {
	fileScanner := bufio.NewScanner(file)
	res := make([]string, 1)
	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		res = append(res, fileScanner.Text())
	}
	return res
}

func main() {
	k := flag.Int("k", -1, "Column to sort")
	n := flag.Bool("n", false, "Number sort")
	r := flag.Bool("r", false, "Reverse sorting")
	u := flag.Bool("u", false, "Without duplicate")
	flag.Parse()
	info := &Info{
		k:        *k,
		n:        *n,
		r:        *r,
		u:        *u,
		fileName: os.Args[len(os.Args)-1],
	}

	file := info.openFile()
	res := info.fileProc(file)
	info.Sorting(&res)
	for _, re := range res {
		if re != "" {
			fmt.Println(re)
		}
	}
}

func (flags *Info) Sorting(slice *[]string) {
	sort.Strings(*slice)

	if flags.r {
		RSorting(*slice)
	}

	if flags.u {
		*slice = USorting(*slice)
	}

	if flags.k > -1 {
		*slice = KSorting(*slice, flags.k)
	}
}

func RSorting(text []string) {
	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		text[i], text[j] = text[j], text[i]
	}
}

func USorting(text []string) []string {
	m := make(map[string]struct{})
	res := make([]string, 0)
	a := struct{}{}
	for _, s := range text {
		_, ok := m[s]
		if !ok {
			m[s] = a
			res = append(res, s)
		}
	}
	return res
}

func KSorting(text []string, k int) []string {
	tmp := make([]string, 0)
	res := make([]string, 0)
	for _, s := range text {
		tmp = append(tmp, strings.Split(s, " ")...)
	}
	fmt.Println(tmp)
	return res
}
