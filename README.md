# yes

Yes, another [ANSI escape code](https://en.wikipedia.org/wiki/ANSI_escape_code) package.
This one differs from the others in that it is the worst one.

## Install

    go get github.com/ofabricio/yes

## Examples

```go
package main

import "github.com/ofabricio/yes"

func main() {

    fmt.Print(yes.AltBufferOn, yes.CursorOn, yes.Cursor11)
    defer fmt.Print(yes.AltBufferOff)

    fmt.Print(yes.FgYellow, "Hello", " ", yes.Underline, "World", yes.Reset)

    bufio.NewReader(os.Stdin).ReadLine() // Press Enter to exit.
}
```

```go
// Print Hello with a red foreground color and reset the style for the next print.
fmt.Print(yes.FgRed, "Hello", yes.Reset)
```

```go
// Join styles into one.
keyword := fmt.Sprint(yes.FgBlue, yes.Underline)

fmt.Print(keyword, "Hello", yes.Reset)
```

```go
fmt.Println(yes.FgRed, "16 colors")
fmt.Println(fmt.Sprintf(yes.Fg256, 2), "256 colors")
fmt.Println(fmt.Sprintf(yes.FgRGB, 200, 90, 30), "RGB colors")
```
