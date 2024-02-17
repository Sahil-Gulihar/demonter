package main

import (
	"io/ioutil"
	"log"
	"strings"
)


func main() {

	for i:=0 ; i<100 ; i++ {
		msg := []byte (strings.Repeat("Hello",1024*1024*300 ))
		err := ioutil.WriteFile("hello", msg , 0644)

		if err != nil{
			log.Fatal(err)
		}
	}
	for i:=0 ; i<100 ; i++ {
		msg := []byte (strings.Repeat("Hello1",1024*1024*300 ))
		err := ioutil.WriteFile("hello", msg , 0644)

		if err != nil{
			log.Fatal(err)
		}
	}

}
