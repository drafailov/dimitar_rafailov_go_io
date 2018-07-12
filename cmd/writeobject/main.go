package main

import (
	"os"
	"fmt"
	"github.com/drafailov/dimitar_rafailov_go_io/goserialize"
)

func main() {

	type address struct {
		City        string
		AddressLine string
		PostCode    string
	}

	addr := address{"Sofia", "Ivan Vazov", "1000"}
	err := goserialize.WriteGob(os.Args[1], addr)
	if err != nil {
		fmt.Println(err)
	}

	var adressRead = new(address)
	err = goserialize.ReadGob(os.Args[1], adressRead)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", adressRead)
	}
}
