package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	f := flag.String("f", "without", "выбрать поля (колонки); формат ввода: 0,1,2,3")
	d := flag.String("d", "\t", "использовать другой разделитель")
	s := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	// Сохраняем индексы запрашиваемых столбцов
	var cols []int
	if *f != "without" {
		colsInString := strings.Split(*f, ",")
		for _, col := range colsInString {
			colInInt, err := strconv.Atoi(col)
			if err != nil {
				panic("некорректный флаг f")
			}
			cols = append(cols, colInInt)
		}
	}

	// Ввод строк
	lines := make([]string, 0)
	fmt.Println("Введите строки для обработки. Чтобы закончить ввод, введите пустую строку.")
	reader := bufio.NewReader(os.Stdin)
	for true {
		input, _ := reader.ReadString('\n')
		if input == "\n" {
			break
		}
		lines = append(lines, input)
	}

	resLines := make([]string, 0)
	for _, line := range lines {
		splitted := strings.Split(line, *d)
		// Если режим без разделителей и Split выдал слайс с одним элементом - пропускаем
		if *s && len(splitted) == 1 {
			continue
		}
		// Если индексы столбцов не указаны, либо разделителей нет - возвращаем всю строку
		if len(cols) == 0 || len(splitted) == 1 {
			resLines = append(resLines, line)
			continue
		}
		// Формируем строку из запрашиваемых столбцов
		newLine := ""
		for _, i := range cols {
			newLine += splitted[i] + *d
		}
		newLine += "\n"
		resLines = append(resLines, newLine)
	}

	// Вывод
	for _, line := range resLines {
		fmt.Print(line)
	}
}
