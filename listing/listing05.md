Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```

Вывод программы объясняется тем, интерфейсная переменная типа error содержит в себе указатель на *customError и не равна nil. Поэтому рекомендуется работать непосредственно с типом error.

Исправить ситуацию можно с помощью type assertion:
```
if err.(*customError) != nil {
	println("error")
	return
}
```
