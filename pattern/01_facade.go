package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Фасад - это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов.
Этот паттерн полезен, когда нужно предоставить урезанный интерфейс к сложной подсистеме.
Пример: чтобы упростить взаимодействие со сложной системой авторизации, можно добавить фасад.
Плюсы:
- Уменьшение зависимости клиента от внутренней структуры
- Предоставляет простой способ взаимодействия с подсистемами
Минусы:
- Повышение сложности из-за добавления еще одного слоя абстракции
- Для часто используемых через фасад операций может наблюдаться ухудшение производительности
*/

type Inputer struct{}

func (inp *Inputer) Input() (login, password string) {
	fmt.Println("Введите логин:")
	fmt.Scanln(&login)
	fmt.Println("Введите пароль:")
	fmt.Scanln(&password)
	return
}

type Validator struct{}

func (val *Validator) Validate(login, password string) error {
	if len(login) <= 5 {
		return errors.New("ошибка валидации: логин должен быть > 5")
	}
	if len(password) <= 8 {
		return errors.New("ошибка валидации: пароль должен быть > 8")
	}
	return nil
}

type Database struct {
	login    string
	password string
}

func (db *Database) IsDataCorrect(login, password string) bool {
	if db.login == login && db.password == password {
		return true
	}
	return false
}

type Facade struct {
	inp *Inputer
	val *Validator
	db *Database
}

func NewFacade(corrLogin, corrPassword string) *Facade {
	return &Facade{
		inp: &Inputer{},
		val: &Validator{},
		db: &Database{login: corrLogin, password: corrPassword},
	}
}

func (f *Facade) CheckUser() bool {
	login, password := f.inp.Input()
	err := f.val.Validate(login, password)
	if err != nil {
		panic(err)
	}
	return f.db.IsDataCorrect(login, password)
}

func main() {
	facade := NewFacade("correctLogin", "correctPassword")
	fmt.Println(facade.CheckUser())
}
