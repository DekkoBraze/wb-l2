package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Фабричный метод - это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе,
позволяя подклассам изменять тип создаваемых объектов. Выгоден, т к при добавлении новых типов объектов в код не придется переписывать всю программу.
Паттерн полезен, когда неизвестно заранее, какие типы объектов необходимо создать
Пример: легкорасширяемая системы создания объектов, н-р, регистрация транспорта в системе
Плюсы:
- Легко расширять
- Более универсальный подход к созданию объектов
Минусы:
- Для расширения необходимо создавать новые структуры
*/

// Продукт
type Vehicle interface {
	GetInfo()
}

// Мотоцикл
type Bike struct {
	Name string
}

func (v *Bike) GetInfo() {
	fmt.Printf("This is a bike - %v\n", v.Name)
}

// Автомобиль
type Car struct {
	Name string
}

func (v *Car) GetInfo() {
	fmt.Printf("This is a car - %v\n", v.Name)
}

// Фабрика
type Factory interface {
	Create()
}

// Конкретная фабрика, создающая мотоциклы
type BikeFactory struct{}

func (f *BikeFactory) Create() Vehicle {
	return &Bike{Name: "Урал"}
}

// Конкретная фабрика, создающая автомобили
type CarFactory struct{}

func (f *CarFactory) Create() Vehicle {
	return &Car{Name: "Газель"}
}

// Пример использования
func FactoryMethodExample() {
	factory := BikeFactory{}
	bike := factory.Create()
	bike.GetInfo()
}
