package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func cd(folder string, currentDir string) string{
	if err := os.Chdir(folder); err != nil {
		fmt.Println(err)
	}
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return currentDir
}

func pwd() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(currentDir)
}

func echo(arg string) {
	fmt.Println(arg)
}

func kill(pid int) {
	syscall.Kill(pid, syscall.SIGTERM)
}

func ps() {
	out, err := exec.Command("/bin/sh", "-c", "ps").Output()
	if err != nil {
		fmt.Println(err)
	}
	strOut := string(out)
	fmt.Println(strOut[:len(strOut)-1])
}

func nc(ip string, port string, message string) {
	tcpServer, err := net.ResolveTCPAddr("tcp", ip+":"+port)
	if err != nil {
		fmt.Println(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpServer)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Получаем текущую директорию и отображаем её
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Добро пожаловать в shell! Чтобы выйти, введите /q.")
	fmt.Printf(currentDir + "$ ")
	for {
		input, _ := reader.ReadString('\n')
		if input == "/q\n" {
			break
		}
		splittedInput := strings.Split(input, " ")
		switch splittedInput[0] {
		case "cd":
			// Удаляем знак переноса строки из аргумента
			arg := splittedInput[1][:len(splittedInput[1])-1]
			currentDir = cd(arg, currentDir)
		case "pwd\n":
			pwd()
		case "echo":
			arg := splittedInput[1][:len(splittedInput[1])-1]
			echo(arg)
		case "kill":
			arg, err := strconv.Atoi(splittedInput[1])
			if err != nil {
				fmt.Println("Некорректные аргументы команды.")
				break
			}
			kill(arg)
		case "ps\n":
			ps()
		case "nc":
			if len(splittedInput) != 4 {
				fmt.Println("Некорректные аргументы команды.")
				break
			}
			nc(splittedInput[1], splittedInput[2], splittedInput[3][:len(splittedInput[3])-1])
		default:
			if len(splittedInput) != 2 {
				fmt.Println("Некорректные аргументы команды.")
				break
			}
			fmt.Println("Неизвестная команда.")
		}
		fmt.Printf(currentDir + "$ ")
	}
	fmt.Println("До встречи!")
}
