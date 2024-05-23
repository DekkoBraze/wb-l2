package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Цепочка вызовов - это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.
Паттерн полезен, когда существует несколько обработчиков, которые могут обработать объект, либо если набор объектов задается динамически.
Пример: обработка значений разных типов или форматов при условии, что заранее он тип неизвестен
Плюсы:
- Ослабление связанности объектов
- В цепочку легко добавить новый обработчик
Минусы:
- Никто не гарантирует, что запрос будет обработан
*/

// Обработчик
type Handler interface {
	SetNext(Handler)
	Handle(interface{})
}

// Базовый обработчик, который внедряется в конкретные обработчики
type BaseHandler struct {
	Next Handler
}

func (h *BaseHandler)SetNext(nextHandler Handler) {
	h.Next = nextHandler
}

// Обработчик строк
type StringHandler struct {
	BaseHandler
}

func (h *StringHandler)Handle(variable interface{}) {
	newVar, ok := variable.(string)
	if !ok && h.Next != nil {
		h.Next.Handle(variable)
		return
	}
	fmt.Printf("%v - это строка!\n", newVar)
}

// Обработчик чисел
type IntHandler struct {
	BaseHandler
}

func (h *IntHandler)Handle(variable interface{}) {
	newVar, ok := variable.(int)
	if !ok && h.Next != nil {
		h.Next.Handle(variable)
		return
	}
	fmt.Printf("%v - это число!\n", newVar)
}

// Обработчик мап
type MapStringIntHandler struct {
	BaseHandler
}

func (h *MapStringIntHandler)Handle(variable interface{}) {
	newVar, ok := variable.(map[string]int)
	if !ok && h.Next != nil {
		h.Next.Handle(variable)
		return
	}
	fmt.Printf("%v - это мапа!\n", newVar)
}

// Пример работы
func ChainExample() {
	intHandler := IntHandler{}
	stringHandler := StringHandler{}
	mapHandler := MapStringIntHandler{}
	intHandler.Next = &stringHandler
	stringHandler.Next = &mapHandler
	intHandler.Handle(map[string]int{})
}
