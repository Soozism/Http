package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// first way to show body
	b := make([]byte, 99999)
	resp.Body.Read(b)
	fmt.Println(string(b))
	
	// seccond way to show body
	io.Copy(os.Stdout, resp.Body)
}
