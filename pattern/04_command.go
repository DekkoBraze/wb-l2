package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Комманда - это поведенческий паттерн проектирования, который превращает запросы в объекты, позволяя передавать их как аргументы при вызове методов,
ставить запросы в очередь, логировать их, а также поддерживать отмену операций.
Паттерн полезен, когда необходимо осуществить поддержку передачи действия аргументом, отмены операций и очереди запросов.
Пример: операции с сервером
Плюсы:
- Позволяет реализовать операцию отмены команды
- Позволяет реализовать отложенный запуск
- Убирает зависимость команды от объекта (можно передавать команду в качестве аргумента)
- Позволяет создавать наборы команд с различными параметрами
Минусы:
- Усложняет код
*/

// Комманда
type Command interface {
	Execute()
}

// Сервер (receiver)
type Server struct {
	Working bool
	Ip string
	Port string
	Mode string
}

// Включение сервера с параметрами
func (s *Server)On(ip, port, mode string) {
	if s.Working {
		fmt.Printf("Сервер уже был запущен по адресу %v:%v в режиме %v\n", ip, port, mode)
		return
	}
	s.Working = true
	s.Ip = ip
	s.Port = port
	s.Mode = mode
	fmt.Printf("Сервер запущен по адресу %v:%v в режиме %v\n", ip, port, mode)
}

// Выключение сервера
func (s *Server)Off() {
	s.Ip = ""
	s.Port = ""
	s.Mode = ""
	s.Working = false
	fmt.Println("Сервер выключен")
}

// Команда включения сервера
type CommandServerOn struct {
	Server Server
	Ip string
	Port string
	Mode string
}

func (c *CommandServerOn)Execute() {
	c.Server.On(c.Ip, c.Port, c.Mode)
}

// Команда выключения сервера
type CommandServerOff struct {
	Server Server
}

func (c *CommandServerOff)Execute() {
	c.Server.Off()
}

// Вызыватель команды
type Invoker struct {
	command Command
}

func (i *Invoker)SetCommand(c Command) {
	i.command = c
}

func (i *Invoker)ExecuteCommand() {
	i.command.Execute()
}

// Пример работы
func CommandExample() {
	server := Server{}
	invoker := Invoker{}
	invoker.SetCommand(&CommandServerOn{Server: server, Ip: "128.163.15.20", Port: "1234", Mode: "low-energy"})
	invoker.ExecuteCommand()
	invoker.SetCommand(&CommandServerOff{})
	invoker.ExecuteCommand()
}