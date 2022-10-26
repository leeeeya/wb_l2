package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type wget struct {
	addr string
	html []byte
}

// Устанавливает адрес в структуру
func (w *wget) SetAddr(str string) {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		log.Fatal("incorrect URL:", err)
	}
	w.addr = str
}

/*
....GetBody()
....Возвращает html разметку страницы
*/
func (w *wget) GetBody() string {
	req, err := http.Get(w.addr)
	if err != nil {
		log.Fatal("Get error:", err)
	}
	defer req.Body.Close()
	w.html, err = io.ReadAll(req.Body)
	if err != nil {
		log.Fatal("Read error:", err)
	}
	//w.html = append(w.html, b...)
	return string(w.html)
}

// Записывает html в структуру
func (w *wget) WriteFile() error {
	file, err := os.Create("index.html")
	if err != nil {
		return err
	}
	_, err = file.WriteString(string(w.html))
	if err != nil {
		return err
	}
	log.Println("Complete: index.html")
	return nil
}

func main() {
	wgetStruct := new(wget)
	wgetStruct.SetAddr(os.Args[len(os.Args)-1])
	wgetStruct.GetBody()
	err := wgetStruct.WriteFile()
	if err != nil {
		log.Fatal(err)
	}
}
