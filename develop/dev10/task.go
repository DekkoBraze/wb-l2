package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeoutFlag := flag.Uint("timeout", 10, "указать таймаут подключения")
	flag.Parse()

	if flag.NArg() < 2 {
		panic(errors.New("ip или порт не указаны"))
	}

	ip := flag.Arg(0)
	port := flag.Arg(1)
	timeout := time.Duration(*timeoutFlag) * time.Second

	// Коннектимся к указанному сокету с таймаутом
	conn, err := net.DialTimeout("tcp", ip+":"+port, timeout)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	readerForSend := bufio.NewReader(os.Stdin)
	readerForReceive := bufio.NewReader(conn)

	wg := sync.WaitGroup{}
	wg.Add(2)

	// Запись в указанный сокет
	go func() {
		for {
			input, err := readerForSend.ReadString('\n')
			// Если 
			if err != nil {
				if err == io.EOF {
					conn.Close()
					wg.Done()
					return
				}
			}
			_, err = conn.Write([]byte(input))
			if err != nil {
				panic(err)
			}
		}
	}()

	// Вывод сообщений из указанного сокета
	go func() {
		for {
			message, err := readerForReceive.ReadString('\n')
			if err != nil {
				fmt.Println("Соединение закрыто.")
				conn.Close()
				wg.Done()
				return
			}
			fmt.Print(message)
		}
	}()

	wg.Wait()
}
