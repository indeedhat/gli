package main

import (
    "fmt"
    "sort"
)

func main() {
    strings := []string{
        "aaaaaaaa",
        "a",
        "Aaaa",
        "aa",
        "aaa",
    }

    sort.Slice(strings, func (i, j int) bool {
        return len(strings[i]) > len(strings[j])
    })

    fmt.Println(strings)
}
