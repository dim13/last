package main

import (
	"fmt"
	"log"

	"github.com/dim13/last"
)

func main() {
	l, err := last.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(l)
}
