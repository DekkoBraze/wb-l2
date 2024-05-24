package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Функция получения итогового массива с учетом контекста и нумерации
func resWithContext(a, b, c int, n bool, filteredLines, allLines []string) []string {
	if a < c {
		a = c
	}
	if b < c {
		b = c
	}

	removedLinesIndexes := make([]int, 0)
	resLines := make([]string, 0)

	for i := 0; i < len(allLines); i++ {
		if len(filteredLines) == 0 {
			break
		}
		if filteredLines[0] == allLines[i] {
			// Определяем нижнюю и верхнюю границу контекста
			infinum := int(math.Max(float64(i-b), 0))
			supremum := int(math.Min(float64(i+a), float64(len(allLines)-1)))
			for j := infinum; j <= supremum; j++ {
				// Если элемент уже добавлен в итоговый список - пропускаем
				if len(removedLinesIndexes) != 0 && removedLinesIndexes[len(removedLinesIndexes)-1] >= j {
					continue
				}
				// Если попадается элемент из списка отфильтрованных строк - расширяем верхнюю границу с его учетом
				if len(filteredLines) != 0 && filteredLines[0] == allLines[j] {
					filteredLines = filteredLines[1:]
					supremum = int(math.Min(float64(j+a), float64(len(allLines)-1)))
				}
				// С нумерацией и без
				if n {
					resLines = append(resLines, fmt.Sprintf("%v. ", j+1)+allLines[j])
				} else {
					resLines = append(resLines, allLines[j])
				}
				removedLinesIndexes = append(removedLinesIndexes, j)
			}
			continue
		}
	}
	return resLines
}

func main() {
	A := flag.Int("A", 0, "печатать +N строк после совпадения")
	B := flag.Int("B", 0, "печатать +N строк до совпадения")
	C := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	c := flag.Bool("c", false, "количество строк")
	i := flag.Bool("i", false, "игнорировать регистр")
	v := flag.Bool("v", false, "вместо совпадения, исключать")
	F := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "напечатать номер строки")

	flag.Parse()

	//Если нет безымянного аргумента, указывающего название файла - ошибка
	if flag.NArg() < 2 {
		panic(errors.New("имя файла или шаблон поиска не указаны"))
	}

	// Считываем название файла и шаблон
	var line string
	var template *regexp.Regexp
	if *F {
		line = flag.Args()[0]
		if *i {
			line = strings.ToLower(line)
		}
	} else {
		var err error
		template, err = regexp.Compile(flag.Args()[0])
		if err != nil {
			panic(err)
		}
	}
	fileName := flag.Args()[1]
	file, err := os.Open(fileName)
	if file == nil {
		panic(errors.New("файл с данным именем не найден"))
	}
	if err != nil {
		panic(errors.New("файл не имеет текстовый формат"))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	filteredLines := make([]string, 0)
	allLines := make([]string, 0)

	for scanner.Scan() {
		newLine := scanner.Text()
		allLines = append(allLines, newLine)
		if *i {
			newLine = strings.ToLower(newLine)
		}
		isMatched := false
		if !*F && template.MatchString(newLine) || *F && line == newLine {
			isMatched = true
		}
		if *v {
			isMatched = !isMatched
		}
		if isMatched {
			filteredLines = append(filteredLines, scanner.Text())
		}
	}

	resLines := resWithContext(*A, *B, *C, *n, filteredLines, allLines)

	for _, v := range resLines {
		fmt.Println(v)
	}
	if *c {
		fmt.Println("Количество подходящих строк: " + fmt.Sprint(len(filteredLines)))
	}
}
