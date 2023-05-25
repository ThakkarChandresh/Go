package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down!", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d\n", url, resp.StatusCode)

		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"

			if err != nil {
				log.Fatal(err)
			} else {
				err = os.WriteFile(file, bodyBytes, 0664)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func main() {
	urls := []string{"https://www.swaminarayangadi.com", "https://www.google.com", "https://www.medium.com", "https://www.oracle.com"}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))
	}

	wg.Wait()
}
