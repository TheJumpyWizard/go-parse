package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "my-sql-parser/internal/parser"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("> ")
        if !scanner.Scan() {
            break
        }
        input := strings.TrimSpace(scanner.Text())
        if input == "exit" {
            break
        }
        p := parser.NewSQLParser(input)
        result, err := p.Parse()
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Printf("%#v\n", result)
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
    }
}

