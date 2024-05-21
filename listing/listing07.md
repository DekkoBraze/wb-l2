Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
2
4
3
5
7
0
0
0
0
0
...
```

Данный вывод объясняется тем, что, когда канал закрывается, при каждом считывании с него данных он начинает вовзвращать значение типа по-умолчанию - в данном случае 0. Когда одна из записывающих горутин заканчивает свою работу и закрывает канал, блок select начинает бесконечно принимать значение по-умолчанию с закрытого канала. 

Решить данную проблему можно с помощью проверки на то, закрыт ли канал и получены ли все значения из него:
```
if _, ok := a; ok {
// Действия если канал не закрыт или у него еще есть значения для передачи
}
```
