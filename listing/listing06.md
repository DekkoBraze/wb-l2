Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
[3 2 3]
```

Слайс s передается по значению, однако, т. к. представляет собой обертку поинтера на массив, первая строчка в функции меняет содержание s. Затем, т. к. вместимость слайса равна 3, функция append создает новый слайс, который присваивает локальной переменной i. Далее все изменения влияют только на локальную переменную i.
