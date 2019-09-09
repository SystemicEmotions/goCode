package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)
func main() {
	resp, _ := http.Get("https://www.google.com/")
	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("HTML:\n\n", string(bytes))

	resp.Body.Close()
}
