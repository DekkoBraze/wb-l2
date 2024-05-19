package main

import (
	"bufio"
	"errors"
	"flag"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Пример команды: go run task.go -r -u -k=1 strings.txt
// -k=0 - первый столбец, отсчет начинается с нуля

func main() {
	k := flag.Int("k", 0, "Указание колонки для сортировки")
	n := flag.Bool("n", false, "Сортировать по числовому значению")
	r := flag.Bool("r", false, "Сортировать в обратном порядке")
	u := flag.Bool("u", false, "Не выводить повторяющиеся строки")

	flag.Parse()

	//Если нет безымянного аргумента, указывающего название файла - ошибка
	if flag.NArg() == 0 {
		panic(errors.New("имя файла не указано"))
	}

	// Считываем название файла и открываем
	fileName := flag.Args()[0]
	file, err := os.Open(fileName)
	if file == nil {
		panic(errors.New("файл с данным именем не найден"))
	}
	if err != nil {
		panic(errors.New("файл не имеет текстовый формат"))
	}
	defer file.Close()

	lines := make([][]string, 0)
	set := make(map[string]int)
	scanner := bufio.NewScanner(file)

	// Если u - удаляем дубликаты
	if *u {
		for scanner.Scan() {
			set[scanner.Text()] += 1
		}
		for line, count := range set {
			if count == 1 {
				lines = append(lines, strings.Split(line, " "))
			}
		}
	} else {
		for scanner.Scan() {
			lines = append(lines, strings.Split(scanner.Text(), " "))
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Задаем функцию Less для сортировки слайса
	sort.SliceStable(lines, func(i, j int) bool {
		indexForSort := *k
		// Если n - сравниваем числа
		if *n {
			num1, err1 := strconv.Atoi(lines[i][indexForSort])
			num2, err2 := strconv.Atoi(lines[j][indexForSort])
			if err1 == nil && err2 == nil {
				if *r {
					return num1 > num2
				}
				return num1 < num2
			}
		}
		// Если попадаются одинаковые элементы - берем последующие индексы
		for indexForSort < len(lines[i])-1 && lines[i][indexForSort] == lines[j][indexForSort] {
			indexForSort++
		}

		comparingElements := []string{lines[i][indexForSort], lines[j][indexForSort]}

		if *r {
			return !sort.StringsAreSorted(comparingElements)
		}
		return sort.StringsAreSorted(comparingElements)
	})

	// Вывод в файл
	output, err := os.Create("outputStrings.txt")
	if err != nil {
		panic(err)
	}
	defer output.Close()

	for i, line := range lines {
		output.WriteString(strings.Join(line, " "))
		if i != len(lines)-1 {
			output.WriteString("\n")
		}
	}
}
