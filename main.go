package main

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func writeFile(filename string, msg []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	err := ioutil.WriteFile(filename, msg, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go writeFile("hello", []byte(strings.Repeat("Hello", 1024*1024*300)), &wg)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go writeFile("hello", []byte(strings.Repeat("Hello1", 1024*1024*300)), &wg)
	}

	wg.Wait()
}
