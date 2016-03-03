package main;

import (
	"fmt"
	"log"
	//"image"
	"io/ioutil"
)

func main() {
	fmt.Println("HI")
	file, err := ioutil.ReadFile("img.jpg")
	
	if err != nil {
		log.Fatalf("Error reading image: %s", err)
	}

	fmt.Println(string(file))
}
