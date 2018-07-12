package goserialize

import (
	"os"
	"encoding/gob"
)
// WriteGob write struct to a file
func WriteGob(filePath string, object interface{}) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

// ReadGob read struct from a file
func ReadGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}