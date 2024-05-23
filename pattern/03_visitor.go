package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Посетитель - это поведенческий паттерн проектирования, который позволяет создавать новые операции, не меняя классы объектов,
над которыми эти операции могут выполняться.
Пример: разные типы сериализации информации о различных типах клиентах банка
Плюсы:
- Добавление функциональности без изменения исходных классов
- Реализация двойной диспетчеризации - вызов метода зависит от вызывающего объекта и переданного аргумента
- Возможность описать свой алгоритм по одной операции для каждого объекта
Минусы:
- Трудно добавлять и модифицировать классы
*/

type Visitor interface {
	VisitPerson(*Person) interface{}
	VisitCompany(*Company) interface{}
}

type Account interface {
	Accept(Visitor) interface{}
}

type VisitorJson struct{}

func (v *VisitorJson) VisitPerson(p *Person) interface{}{
	dict := make(map[string]string)
	dict["FirstName"] = p.FirstName
	dict["LastName"] = p.LastName
	dict["Number"] = p.Number
	return dict
}

func (v *VisitorJson) VisitCompany(c *Company) interface{}{
	dict := make(map[string]string)
	dict["Name"] = c.Name
	dict["RegNumber"] = c.RegNumber
	return dict
}

type VisitorHtml struct{}

func (v *VisitorHtml) VisitPerson(p *Person) interface{}{
	result := "<table><tr><td>Свойство<td><td>Значение</td></tr>";
    result += "<tr><td>Name<td><td>" + p.FirstName + "</td></tr>";
    result += "<tr><td>Number<td><td>" + p.LastName + "</td></tr>";
	result += "<tr><td>Number<td><td>" + p.Number + "</td></tr></table>";
	return result
}

func (v *VisitorHtml) VisitCompany(c *Company) interface{}{
	result := "<table><tr><td>Свойство<td><td>Значение</td></tr>";
    result += "<tr><td>Name<td><td>" + c.Name + "</td></tr>";
    result += "<tr><td>Number<td><td>" + c.RegNumber + "</td></tr></table>";
	return result
}

type Person struct {
	FirstName string
	LastName  string
	Number    string
}

func (p *Person) Accept(v Visitor) (res interface{}){
	res = v.VisitPerson(p)
	return
}

type Company struct {
	Name      string
	RegNumber string
}

func (c *Company) Accept(v Visitor) (res interface{}){
	res = v.VisitCompany(c)
	return
}

func main() {
	pers := Person{FirstName: "Foo", LastName: "Barovich", Number: "12345"}
	comp := Company{Name: "Google", RegNumber: "12345"}
	resJson := pers.Accept(&VisitorJson{})
	fmt.Println(resJson)
	resHtml := comp.Accept(&VisitorHtml{})
	fmt.Println(resHtml)
}
