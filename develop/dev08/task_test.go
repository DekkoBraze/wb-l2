package main

import (
	"net"
	"os"
	"testing"
)

func TestCd(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	os.Mkdir("testFolder", 0777)
	rightPath := currentDir + "/testFolder"
	defer os.RemoveAll(rightPath)
	cd("testFolder", currentDir)
	currentDir, err = os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if currentDir != rightPath {
		t.Errorf("handler returned unexpected message: got %v want %v",
			currentDir, rightPath)
	}
}

func TestNc(t *testing.T) {
	listen, err := net.Listen("tcp", "localhost"+":"+"8080")
	if err != nil {
		t.Fatal(err)
	}
	defer listen.Close()

	ch := make(chan bool)
	go func() {
		for {
			_, err := listen.Accept()
			if err != nil {
				ch <- false
			}
			ch <- true
			break
		}
	}()
	
	nc("localhost", "8080", "Hello!")
	ans := <-ch
	if !ans {
		t.Error("Ошибка получения сообщения")
	}
}
