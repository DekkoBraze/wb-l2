package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var errIncorrect = errors.New("передана некорректная строка")

// Unpack : Распаковка строки
func Unpack(symbols string) (unpacked string, err error) {
	runes := []rune(symbols)
	newRunes := []rune{}
	var repeatedRune rune
	numbers := []rune{}
	digitIsNext := false
	escapeIsNext := false
	for i, symbol := range runes {
		if !escapeIsNext && unicode.IsDigit(symbol) {
			// Возвращаем ошибку в случае, если число идет первым в переданной строке
			if !digitIsNext {
				return "", errIncorrect
			}
			// Считаем числа, чтобы учитывать многоразрядные случаи
			numbers = append(numbers, symbol)
			// Если следующий символ не цифра, либо он последний - добавляем повторяющиеся символы в новую строку
			if i < len(runes)-1 && !unicode.IsDigit(runes[i+1]) || i == len(runes)-1 {
				digitIsNext = false
				var counter int
				counter, err = strconv.Atoi(string(numbers))
				if err != nil {
					return
				}
				for j := 0; j < counter; j++ {
					newRunes = append(newRunes, repeatedRune)
				}
				numbers = nil
			}
		} else {
			// Если \ - следующий символ не будет считаться числом
			if !escapeIsNext && string(symbol) == `\` {
				escapeIsNext = true
				continue
			}
			if escapeIsNext {
				escapeIsNext = false
			}
			// Если следующий символ - цифра, запоминаем текущий символ и включаем флаг 
			if !digitIsNext && i < len(runes)-1 && unicode.IsDigit(runes[i+1]) {
				repeatedRune = symbol
				digitIsNext = true
				continue
			}
			// Добавляем нечисловые символы в новую строку
			newRunes = append(newRunes, symbol)
		}
	}
	return string(newRunes), err
}

func main() {
	unpacked, err := Unpack(`a\\5b100`)
	if err != nil {
		panic(err)
	}
	fmt.Println(unpacked)
}
