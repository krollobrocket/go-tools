package main

import (
    "fmt"
    "os"
    "net/http"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("You need to supply a domain.")
        return
    }
	requestURL := os.Args[len(os.Args) - 1]
	res, err := http.Head(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Server:", res.Header.Get("Server"))
	fmt.Println("X-Powered-By:", res.Header.Get("X-Powered-By"))
}