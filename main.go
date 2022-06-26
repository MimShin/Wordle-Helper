package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"wordle/wordle"

	log "github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <dictionary-file>\n", path.Base(os.Args[0]))
		return
	}
	d := wordle.NewWordle(os.Args[1])

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\nLet's go\n> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Error("Error reading from stdin")
			continue
		}
		if len(input) <= 1 {
			help()
			continue
		}
		words := strings.Fields(input)

		switch words[0] {
		case "p":
			fmt.Println(d)
		case "w":
			fmt.Printf("%v\n%v\n", d, d.Current().Words)
		case "q":
			fmt.Println("\nBye")
			return
		case "u":
			d.Undo()
			fmt.Println(d)
		case "a":
			fmt.Printf("filtering: word=%s, key=%s\n", words[1], words[2])
			if len(words[1]) != len(words[2]) {
				fmt.Println("invalid input, word and key must have the same length")
				break
			}
			d.Filter(words[1], words[2])
			fmt.Println(d)
		default:
			help()
		}

	}
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("h  (help)")
	fmt.Println("a word key  (try new word)")
	fmt.Println("  *: right char at right pos")
	fmt.Println("  +: right char at wrong pos")
	fmt.Println("  -: wrong char")
	fmt.Println("  e.g.\n  a style *--+-")
	fmt.Println("u  (undo last word)")
	fmt.Println("p  (print history)")
	fmt.Println("w  (print words)")
	fmt.Println("q  (quit)")
}
