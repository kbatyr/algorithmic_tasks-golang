package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printstr(s string) {
	for _, let := range s {
		z01.PrintRune(let)
	}
}

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		args := os.Args[1:]
		for i := range args {
			file, err := os.ReadFile(args[i])
			if err != nil {
				// s := "ERROR: "
				s := "cat: "
				s1 := ": No such file or directory"
				printstr(s)
				printstr(args[i])
				printstr(s1)
				// io.WriteString(os.Stdout, err.Error())
				z01.PrintRune(10)
			}
			printstr(string(file))
		}
	}
}
