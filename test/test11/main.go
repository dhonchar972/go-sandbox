package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	WriteToFile()
	ReadFromFile()
}

func ReadFromFile() {
	data, err := ioutil.ReadFile("test.txt") // Deprecated: As of Go 1.16, this function simply calls os.ReadFile.
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) // hello world!
}

func WriteToFile() {
	data := []byte("some text")
	err := ioutil.WriteFile("text.txt", data, 0777) // Deprecated: As of Go 1.16, this function simply calls os.WriteFile.
	if err != nil {
		panic(err)
	}
	fmt.Println("Writing complited!")
}
