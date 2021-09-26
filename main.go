package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func brainfuckEval(code string) {

	tokens := strings.Split(code, "")
	for _, token := range tokens {

	}

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
		brainfuckEval(args[0])
	} else {
		bytes, err := ioutil.ReadFile(args[0])
		if err != nil {
			panic(err)
		}
		brainfuckEval(string(bytes))
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use: "brainfuck",
		Run: brainFuckMain,
	}

	rootCmd.PersistentFlags().BoolP("eval", "e", false, "execute of brainfuck script")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
