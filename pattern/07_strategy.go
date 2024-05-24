package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Стратегия - это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс.
После чего, алгоритмы можно взаимозаменять прямо во время исполнения программы.
Полезен, когда нужно объединить родственные классы с разным поведением или когда необходимо обеспечить выбор вариантов алгоритмов
и менять их по ходу работы программы
Пример: различающиеся в зависимости от типа данных обработка и вывод данных
Плюсы:
- Инкапсуляция реализации алгоритмов
- Однообразный вызов всех алгоритмов
- Отказ от условных операторов
Минусы:
- Нагрузка кода дополнительными классами
*/

// Стратегия
type Strategy interface {
	Sort([]int) []int
}

// Контекст, который хранит текущую стратегию и вызывает её метод
type Context struct {
	Data            []int
	ContextStrategy Strategy
}

func (c *Context) ExecuteSort() []int {
	return c.ContextStrategy.Sort(c.Data)
}

// Сортировка пузырьком
type BubbleSort struct{}

func (s *BubbleSort) Sort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// Быстрая сортировка
type QuickSort struct{}

func (s *QuickSort) Sort(arr []int) (outArr []int) {
	less := make([]int, 0)
	equal := make([]int, 0)
	greater := make([]int, 0)
	if len(arr) > 1 {
		pivot := arr[0]
		for _, x := range arr {
			if x < pivot {
				less = append(less, x)
			} else if x == pivot {
				equal = append(equal, x)
			} else if x > pivot {
				greater = append(greater, x)
			}
		}
		finalLess := s.Sort(less)
		finalGreater := s.Sort(greater)
		outArr = append(finalLess, equal...)
		outArr = append(outArr, finalGreater...)
		return
	} else {
		return arr
	}
}

// Пример работы
func StrategyExample() {
	data := []int{-5, 1, 10, -50, 3, 4, 7}
	context := Context{Data: data, ContextStrategy: &BubbleSort{}}
	res := context.ExecuteSort()
	fmt.Println(res)
	context.ContextStrategy = &QuickSort{}
	res = context.ExecuteSort()
	fmt.Println(res)
}
