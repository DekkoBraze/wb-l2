package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Строитель - это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.
Пример: создание сложной легкорасширяемой системы обращения к пользователю
Плюсы:
- Позволяет использовать один и тот же код для создания отличающихся объектов
- Инкапсулирует код, реализующий конструирование
- Позволяет структурировать конструирование объекта и более тонко настраивать процесс
Минусы:
- Повышение сложности кода
*/

// Интерфейс билдера
type Builder interface {
	MakeGreeting()
	MakeUsername(string)
	MakePunctuation()
	GetResultMessage() string
}

// Билдер, собирающий сообщение с приветствием
type HelloBuilder struct {
	message string
}

func (b *HelloBuilder)MakeGreeting() {
	b.message += "Hello, "
} 

func (b *HelloBuilder)MakeUsername(name string) {
	b.message += name
} 

func (b *HelloBuilder)MakePunctuation() {
	b.message += "!"
} 

func (b *HelloBuilder)GetResultMessage() string{
	return b.message
} 

// Билдер, собирающий сообщение с прощанием
type GoodbyeBuilder struct {
	message string
}

func (b *GoodbyeBuilder)MakeGreeting() {
	b.message += "Goodbye, "
} 

func (b *GoodbyeBuilder)MakeUsername(name string) {
	b.message += name
} 

func (b *GoodbyeBuilder)MakePunctuation() {
	b.message += "..."
} 

func (b *GoodbyeBuilder)GetResultMessage() string{
	return b.message
} 

// Директор для управления билдерами
type Director struct {
	builder Builder
}

// Директорский метод создания сообщения
func (d *Director)CreateMessage(name string) string{
	d.builder.MakeGreeting()
	d.builder.MakeUsername(name)
	d.builder.MakePunctuation()
	return d.builder.GetResultMessage()
}

// Пример
func BuilderExample() {
	b1 := HelloBuilder{}
	b2 := GoodbyeBuilder{}
	dir := Director{}
	dir.builder = &b1
	res1 := dir.CreateMessage("CoolUsername")
	fmt.Println(res1)
	dir.builder = &b2
	res2 := dir.CreateMessage("VeryCoolUsername")
	fmt.Println(res2)
}
