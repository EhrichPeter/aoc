package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(file), "\n")

	fmt.Println(rows)

}
