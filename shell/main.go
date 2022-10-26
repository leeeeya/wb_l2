package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

/*
	Shell утилита, содержит команды - cd/pwd/echo/kill/ps
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	for { // Цикл программы
		cur, _ := user.Current()
		fmt.Print(cur.Name+"@", "-->  ")
		command, err := reader.ReadString('\n') // Чтение программы
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		err = Processing(command) // Выполнение
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
	}
}

func Processing(str string) (err error) {
	non := strings.Trim(str, "\n")      // Убираю перенос строки
	commands := strings.Split(non, ";") // Разделяю по знаку ';' т.к. может быть последовательнось комманд
	for _, command := range commands {
		com := strings.Trim(command, " ")
		args := strings.Split(com, " ")
		switch args[0] { // проверки для cd/exit
		case "cd":
			if len(args) < 2 {
				fmt.Println("path error не написал")
				return nil
			}
			continue
		case "exit":
			fmt.Println("Terminated")
			os.Exit(0)
		}
		cmd := exec.Command(args[0], args[1:]...) // Заполнеие полей cmd структуры
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err = cmd.Run() // Непосредственно выполнеие программы
	}
	return err
}
