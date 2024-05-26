package main

import (
	"errors"
	"flag"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		panic(errors.New("ссылка на ресурс не указана"))
	}

	link := flag.Arg(0)
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}

	// Создаем папку с именем сайта
	splittedLink := strings.Split(link, "//")
	dirName := strings.Split(splittedLink[1], "/")[0]
	err = os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Сохраняем body из респонса
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Создаем файл index.html в новой папке и сохраняем туда body
	f, err := os.Create(dirName + "/index.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(string(bytes))
}
