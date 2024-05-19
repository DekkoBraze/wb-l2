package main

import (
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	os.Args = append(os.Args, "-k=2")
	os.Args = append(os.Args, "-n")
	os.Args = append(os.Args, "-r")
	os.Args = append(os.Args, "strings.txt")

	main()
}
