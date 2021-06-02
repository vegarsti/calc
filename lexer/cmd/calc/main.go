package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/vegarsti/calc/lexer"
	"github.com/vegarsti/calc/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is a calculator!\n", user.Username)
	Start(os.Stdin, os.Stdout)
}
