// Package gofile contains functions to open files and write in it
package gofile

import (
	"os"
	"io"
	"bufio"
	"log"
	"bytes"
	"fmt"
)

// OpenFile returns open file in the given path.
// with write permission and append or create if the file doesn't exist.
func OpenFile(path string) (file *os.File, err error) {
	fmt.Print(path)
	file, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return
	}
	return
}

//Scanner reads content from the provided Reader till it reaches the passed stop word
// and the read content is returned as slice of bytes.

func scaner(r io.Reader, endOfString string) (p []byte) {
	scan := bufio.NewScanner(r)
	exitString := []byte(endOfString)
	for {
		scan.Scan()
		buf := scan.Bytes()
		if err := scan.Err(); err != nil {
			log.Fatalf("Error reading: %v\n", err)
		}

		if bytes.Equal(buf, exitString) {
			break
		} else {
			buf = append(buf, '\n')
			p = append(p, buf...)
		}
	}

	return
}

// BreakableConsoleScanner returns a slice of data bytes from the console
func BreakableConsoleScanner(exitString string) []byte {
	return scaner(os.Stdin, exitString)
}
