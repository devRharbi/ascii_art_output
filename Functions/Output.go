package ascii

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	size     int
	nameFlag string
)

func CheckArgsAndFlag(arg []string) []string {
	flag.CommandLine.SetOutput(io.Discard)
	flag.Usage = func() {
		MessageErrors()
	}
	flag.StringVar(&nameFlag, "output", "", "you have add like this")
	flag.Parse()
	args := flag.Args()
	size = len(nameFlag)
	if (size == 0 && strings.HasPrefix(arg[1], "--output=")) || (size != 0 && !strings.HasPrefix(arg[1], "--output=")) {
		MessageErrors()
		os.Exit(0)
	}
	return args
}

// this accepte  text and change to byte and create file to save data
func SaveInFile(textargs string) {
	text := []byte{}
	for i := range textargs {
		text = append(text, byte(textargs[i]))
	}
	if nameFlag != "" {
		if !strings.HasSuffix(nameFlag, ".txt") {
			MessageErrors()
			os.Exit(0)
		}
		errs := os.WriteFile(nameFlag, text, 0o644)
		if errs != nil {
			os.Exit(0)
		}
	}
}

func PrintResult(sl [][]string, s string) {
	result := PrintAsciArt(sl, s)
	if size == 0 {
		fmt.Print(result)
	} else {
		SaveInFile(result)
	}
}
