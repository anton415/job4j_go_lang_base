package tracker

import (
	"bufio"
	"fmt"
	"os"
)

func NewConsoleInput() *ConsoleInput {
	return &ConsoleInput{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

type ConsoleInput struct {
	scanner *bufio.Scanner
}

func (c *ConsoleInput) Get() string {
	if c.scanner == nil {
		c.scanner = bufio.NewScanner(os.Stdin)
	}
	c.scanner.Scan()
	return c.scanner.Text()
}

type ConsoleOutput struct{}

func (c ConsoleOutput) Out(text string) {
	fmt.Println(text)
}
