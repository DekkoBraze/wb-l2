package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Тип, с помощью которого будем сортировать string
type sortRunes []rune

// Навешиваем методы, чтобы он соответствовал интерфейсу sort.Interface
func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortRunes) Len() int           { return len(s) }

// SortString : Сортировка строки
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// FindAnagrams : Нахождение анаграм
func FindAnagrams(arr []string) *map[string][]string {
	set := make(map[string][]string)
	// wordsCounter нужен для удаления дубликатов
	wordsCounter := make(map[string]int)

	for _, el := range arr {
		el = strings.ToLower(el)

		wordsCounter[el]++
		if wordsCounter[el] > 1 {
			continue
		}

		// Присваиваем отсортированную строку как ключ мапы
		sortedEl := SortString(el)
		set[sortedEl] = append(set[sortedEl], el)
	}

	finalSet := make(map[string][]string)

	for k, val := range set {
		// Удаляем множества из одного элемента
		if len(val) <= 1 {
			delete(set, k)
			continue
		}

		// Делаем ключом в новой мапе первый встреченный элемент множества и присваиваем отсортированное множество
		firstVal := val[0]
		sort.Strings(val)
		finalSet[firstVal] = val
	}

	return &finalSet
}

// Пример работы
func main() {
	initArr := [8]string{"пятак", "листок", "пятка", "тяпка", "столик", "СЛИТОК", "лишний", "столик"}

	pSet := FindAnagrams(initArr[:])

	fmt.Println(*pSet)
}
