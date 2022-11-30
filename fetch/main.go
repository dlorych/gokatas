// Fetch prints content found at URLs defined as command line arguments. Adapted
// from github.com/adonovan/gopl.io/tree/master/ch1/fetch.
//
// Level: beginner
// Topics: web client, net/http
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func print(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s", data)

	return nil
}

func main() {

	for _, url := range os.Args[1:] {

		err := print(url)
		if err != nil {
			log.Print(err)
			continue
		}
	}
}
