package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Состояние - Это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.
Извне создаётся впечатление, что изменился класс объекта.
Полезен, когда поведение объекта может меняться по ходу работы программы.
Пример: простое изменение вывода информации о сервере в зависимости от того, работает он или нет
Плюсы:
- Избавляет от большого количества условных операторов
- Концентрирует код, связанный с конкретным состоянием, в одном месте
Минусы:
- Усложнение кода
*/

// Объект сервера, хранящий состояние (контекст)
type ServerObject struct {
	CurrentState State
}

func (s *ServerObject) GetServerInfo() {
	s.CurrentState.GetInfo()
}

// Состояние
type State interface {
	GetInfo()
}

// Состояние работающего сервера
type WorkingState struct{}

func (s *WorkingState) GetInfo() {
	fmt.Println("Сервер работает!")
}

// Состояние остановленного сервера
type StoppedState struct{}

func (s *StoppedState) GetInfo() {
	fmt.Println("Сервер остановлен!")
}

// Пример работы
func main() {
	server := ServerObject{CurrentState: &WorkingState{}}
	server.GetServerInfo()
	server.CurrentState = &StoppedState{}
	server.GetServerInfo()
}
