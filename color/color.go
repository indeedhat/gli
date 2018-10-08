package color

import "fmt"


// wrap the given text in the color for shell output
func Wrap(text, color string) string {
    return fmt.Sprintf("\033[%sm%s\033[%sm", color, text, None)
}