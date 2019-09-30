package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello")

	file, err := os.Open("./datafile")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println("Scanner read " + scanner.Text())
	}

	if err2 := scanner.Err(); err2 != nil {
		log.Fatal(err2)
	}
}
