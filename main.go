package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	b := make([]byte, 99999)
	resp.Body.Read(b)
	fmt.Println(string(b))
}
