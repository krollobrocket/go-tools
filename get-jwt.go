package main

import (
  "os"
  "fmt"
  "strings"
  "flag"
  "github.com/golang-jwt/jwt"
)

func main() {
    algo := flag.String("algo", jwt.SigningMethodHS512.Alg(), "Algorithm to use.")
    flag.Parse()
    if len(os.Args) < 2 {
      fmt.Println("You need to supply a key.")
      return
    }
    key := []byte(strings.TrimSpace(os.Args[len(os.Args) - 1]))
    var selectedAlgo jwt.SigningMethod = jwt.SigningMethodHS512
    switch *algo {
        case "HS256": selectedAlgo = jwt.SigningMethodHS256
        case "HS384": selectedAlgo = jwt.SigningMethodHS384
        case "HS512": selectedAlgo = jwt.SigningMethodHS512
    }
    token := jwt.NewWithClaims(selectedAlgo, jwt.MapClaims{
        // "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
    })
    ss, err := token.SignedString(key)
    if err != nil {
        fmt.Println("Failed to create signed key %v.", err)
    }
    fmt.Printf("Signed Key created (%s) %s", selectedAlgo.Alg(), ss)
    return
}