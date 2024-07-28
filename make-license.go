package main

import (
    "fmt"
    "encoding/json"
    "encoding/base64"
    "os"
    "flag"
)

func main() {
    key := flag.String("key", "", "The key")
    domain := flag.String("domain", "", "The domain")
    start := flag.String("start-date", "", "The start date")
    end := flag.String("end-date", "", "The end date")
    flag.Parse()
    if len(os.Args) < 6 {
        fmt.Println("You need to supply 6 arguments.")
        return
    }
    if len(*key) < 5 {
        fmt.Println("Licence Key needs to be 32 characters long.")
        return
    }
    content := map[string]interface{}{
		"key": *key,
		"domain": *domain,
		"start": *start,
		"end": *end,
	}
	jsonData, err := json.Marshal(content)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
    licenseData := base64.StdEncoding.EncodeToString([]byte(jsonData))
    os.WriteFile(os.Args[len(os.Args) - 1], []byte(licenseData), 0666)
}