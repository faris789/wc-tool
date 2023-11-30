package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func readfile(flname string) ([]byte, error) {
	b, err := os.ReadFile(flname)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func countLines(s []byte, r rune) int {
	count := 0
	for _, ch := range string(s) {
		if ch == r {
			count = count + 1
		}

	}
	count = count + 1
	return count
}

func countWords(s string) int {
	return len(strings.Fields(s))
}

func printHelp() {
	fmt.Println(`		 
	Help for wc command flags. 
	wc <file-name> or wc -<flag> <file-name>

	-c      The number of bytes in each input file is written to the standard output. 
	-l      The number of lines in each input file is written to the standard output.	
	-m      The number of characters in each input file is written to the standard output. 
	-w      The number of words in each input file is written to the standard output.
	`)
	return
}

var wcflag string
var flname string
var bfile []byte

func main() {

	if len(os.Args) == 1 {
		printHelp()

	} else if len(os.Args) == 2 {
		flname = os.Args[1]
		if flname == "--version" {
			fmt.Println("0.0.1")
		} else {
			bfile, err := readfile(flname)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%d\t%d\t%d\t%s\n", countLines(bfile, '\n'), countWords(string(bfile)), len(bfile), flname)
		}

	} else if len(os.Args) == 3 {
		wcflag = os.Args[1]
		flname = os.Args[2]

		bfile, err := readfile(flname)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch wcflag {
		case "-c":
			fmt.Println("byte count:", len(bfile))
		case "-m":
			fmt.Println("character count:", utf8.RuneCountInString(string(bfile)))
		case "-l":
			fmt.Println("line count:", countLines(bfile, '\n'))
		case "-w":
			fmt.Println("word count:", countWords(string(bfile)))
		default:
			printHelp()
		}

	}
}
