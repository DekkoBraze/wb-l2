Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!
```
Вывод программы объясняется тем, что range по каналу считывает данные из него до тех пор, пока канал не будет закрыт. Именно поэтому и случается дедлок после считывания значений - мы не закрываем канал в пишущей горутине. 

Пофиксить можно, закрыв канал в конце работы пишущей горутины:
```
go func() {
for i := 0; i < 10; i++ {
	ch <- i
	}
close(ch)
}()
```

