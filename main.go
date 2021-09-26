package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func bracketSearch(tokens []string) map[int]int {

	bracketMap := map[int]int{}
	bracketLeftArray := []int{}
	count := 0
	for index, token := range tokens {
		if token == "[" {
			bracketLeftArray = append(bracketLeftArray, index)
			count++
		} else if token == "]" {
			count--
			bracketMap[index] = bracketLeftArray[count]
			bracketMap[bracketLeftArray[count]] = index
			bracketLeftArray = bracketLeftArray[0 : len(bracketLeftArray)-1]
		}
	}

	if count != 0 {
		log.Println("The number of [does not match")
	}

	return bracketMap
}

func brainfuckEval(code string) {
	tape := make([]int, 1000)
	pointer := 0
	index := 0
	tokens := strings.Split(code, "")
	bracketMap := bracketSearch(tokens)
	for index < len(tokens) {
		token := tokens[index]
		switch token {
		case ">":
			pointer++
		case "<":
			pointer--
		case "+":
			tape[pointer]++
		case "-":
			tape[pointer]--
		case ".":
			fmt.Print(string(rune(tape[pointer])))
		case ",":
			var i string
			fmt.Scan(i)
			input := []byte(i)[0]
			tape[pointer] = int(input)
		case "[":
			if tape[pointer] == 0 {
				index = bracketMap[index]
			}
		case "]":
			if tape[pointer] != 0 {
				index = bracketMap[index]
			}
		default:
			continue
		}
		index++
	}
}

func newLineDelete(bytes []byte) []byte {

	result := []byte{}
	for _, b := range bytes {
		if b == 10 {
			continue
		} else if b == 13 {
			continue
		} else {
			result = append(result, b)
		}
	}
	return result
}

func brainFuckMain(c *cobra.Command, args []string) {
	if len(args) == 0 {
		c.Help()
		os.Exit(1)
	}

	executeFlag, err := c.PersistentFlags().GetBool("eval")
	if err != nil {
		c.Help()
		os.Exit(1)
	}

	if executeFlag {
		code := string(newLineDelete([]byte(args[0])))
		brainfuckEval(code)
	} else {
		bytes, err := ioutil.ReadFile(args[0])
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		brainfuckEval(string(newLineDelete(bytes)))
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use: "brainfuck <filename>",
		Run: brainFuckMain,
	}

	rootCmd.PersistentFlags().BoolP("eval", "e", false, "execute of brainfuck script")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
