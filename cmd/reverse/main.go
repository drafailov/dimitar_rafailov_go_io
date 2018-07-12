package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"github.com/drafailov/dimitar_rafailov_go_io/goreverse"
	"log"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Print("Filename is missing!")

		return
	}
	s, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading the file", os.Args[1], err)
	}
	fmt.Print(goreverse.Reverse(string(s)))

}
