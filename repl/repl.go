package repl

import (
	"asteroid/lexer"
	"asteroid/token"
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// Gracefully handle the exit of the REPL
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("Exiting REPL, Goodbye!")
		os.Exit(0)
	}()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "exit" {
			// send a signal to the channel to exit the REPL
			sigChan <- os.Interrupt
			continue
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; {
			fmt.Printf("%+v\n", tok)
			tok = l.NextToken()
		}
	}
}
