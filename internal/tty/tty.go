package tty

import (
    "fmt"
    "os"
    "sync"
    "encoding/json"

    "github.com/fatih/color"
    "github.com/TylerBrock/colorjson"
    "github.com/charmbracelet/x/ansi"
    "github.com/charmbracelet/x/term"
)

var isTTY = sync.OnceValue(func() bool {
    return term.IsTerminal(os.Stdout.Fd())
})

func Printjson(b []byte) {
    var obj map[string]interface{}
    json.Unmarshal(b, &obj)
    formatter := colorjson.NewFormatter()
    formatter.BoolColor = color.New(color.FgWhite)
    formatter.KeyColor = color.New(color.FgBlue)
    formatter.Indent = 2
    str, _ := formatter.Marshal(obj)
    Println(string(str))
}
// Println handles println, striping ansi sequences if stdout is not a tty.
func Println(s string) {
    if isTTY() {
        fmt.Println(s)
        return
    }
    fmt.Println(ansi.Strip(s))
}
