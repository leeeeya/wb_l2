package grepSupport

import (
	"flag"
	"os"
)

type Flags struct {
	a            int
	b            int
	cSmall       bool
	cBig         int
	i            bool
	v            bool
	f            bool
	n            bool
	DesiredValue string
	FileName     string
}

func (flagStruct *Flags) Input() {
	a := flag.Int("A", -1, "\"after\" печатать +N строк после совпадения\n")
	b := flag.Int("B", -1, "\"before\" печатать +N строк до совпадения\n")
	cBig := flag.Int("C", -1, "\"context\" (A+B) печатать ±N строк вокруг совпадения\n")
	cSmall := flag.Bool("c", false, "\"count\" количество строк\n")
	i := flag.Bool("i", false, "\"ignore-case\" игнорировать регистр")
	v := flag.Bool("v", false, "\"invert\" вместо совпадения, исключать")
	f := flag.Bool("f", false, "\"fixed\" очное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "\"line-num\" напечатать номер строки")
	flag.Parse()

	//for qwe, arg := range os.Args {
	//	if arg != "-A" && arg != *b && arg != "-c" && arg != "-C" && arg != "-i" && arg != "-v" && arg != "-f" && arg != "-n" && qwe != 0 {
	//		flagStruct.FileName = arg
	//		fmt.Println(arg, qwe, "!!!")
	//		break
	//	}
	//}

	flagStruct.a = *a
	flagStruct.b = *b
	flagStruct.cSmall = *cSmall
	flagStruct.cBig = *cBig
	flagStruct.i = *i
	flagStruct.v = *v
	flagStruct.f = *f
	flagStruct.n = *n
	flagStruct.DesiredValue = os.Args[len(os.Args)-2]
	flagStruct.FileName = os.Args[len(os.Args)-1]
}
