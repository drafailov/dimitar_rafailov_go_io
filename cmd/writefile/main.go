package main

import (
	"fmt"
	"os"
	"github.com/drafailov/dimitar_rafailov_go_io/gofile"
	"log"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Print("Filename is missing!")
		return
	}
	file, err := gofile.OpenFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening the file", os.Args[1], err)
	}

	fmt.Println("Write text to be written into a file. Enter line with only one point for finish: ")
	fmt.Print("-> ")

	bytes, err := file.Write(gofile.BreakableConsoleScanner("."))
	if err != nil {
		log.Fatalf("Error writing to file: ", err)
	}
	file.Close()
	println(bytes)
}
