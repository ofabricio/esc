package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ofabricio/yes"
)

func main() {

	fmt.Print(yes.AltBufferOn, yes.CursorOn, yes.Cursor11)
	defer fmt.Print(yes.AltBufferOff)

	fmt.Print(yes.FgYellow, "Hello", " ", yes.Underline, "World", yes.Reset)

	bufio.NewReader(os.Stdin).ReadLine() // Press Enter to exit.
}
