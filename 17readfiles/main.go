package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in the file - mylcofile.txt"
	file, err := os.Create("./mylcofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is:", length)

	defer file.Close()
	readFile("./mylcofile.txt")
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	content := string(dataByte)
	fmt.Println("The content of the file is:", content)
}
