package main

import (
    "fmt"
    "encoding/base64"
    "log"
    "os"
)

func main() {
    content, err := os.ReadFile(os.Args[len(os.Args) - 1])
    if err != nil {
        log.Fatal(err)
    }
    var str []byte
    str, err = base64.StdEncoding.DecodeString(string(content))
    fmt.Println(string(str))
}