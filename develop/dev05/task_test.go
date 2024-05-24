package main

import (
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	os.Args = append(os.Args, "-C=3")
	os.Args = append(os.Args, "-n")
	os.Args = append(os.Args, "-F")
	os.Args = append(os.Args, "-c")
	os.Args = append(os.Args, "correct")
	os.Args = append(os.Args, "lines.txt")

	main()
}
