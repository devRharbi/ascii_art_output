package main

import (
	"os"
	"strings"

	ascii "omar/Functions"
)

// size     int
var style string

// nameFlag string

func main() {
	args := os.Args
	if len(args) <= 1  {
		ascii.MessageErrors()
		os.Exit(0)
	}
	test := ascii.CheckArgsAndFlag(args)
	if len(test) == 0 {
		ascii.MessageErrors()
		os.Exit(0)
	}
	args1 := test[0]
	for i := range args1 {
		if args1[i] < 32 || args1[i] > 126 {
			ascii.MessageErrors()
		}
	}
	if len(test) == 1 {
		style = "standard.txt"
		slie := ascii.WriteTextFileAscii(style)
		ascii.PrintResult(slie, args1)
	} else {
		style := test[1]
		removeTXt := strings.TrimSuffix(style, ".txt")
		if removeTXt == "shadow" || removeTXt == "standard" || removeTXt == "thinkertoy" || removeTXt == "smile" {
			slie := ascii.WriteTextFileAscii(removeTXt + ".txt")
			ascii.PrintResult(slie, args1)
		} else {
			ascii.MessageErrors()
		}
	}
}
