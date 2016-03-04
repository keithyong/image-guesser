package main;

import (
	"fmt"
    "image"
    "image/jpeg"
	"os"
)

func main() {
	file, err := os.Open("img.jpg")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    img, _, err := image.Decode(file)
    if err != nil {
        panic(err)
    }
    
    fmt.Println(img.Bounds())
    fmt.Println(img.At(10, 10))
    newImg, _ := os.Create("new.jpg")
    defer newImg.Close()
    
    jpeg.Encode(newImg, img, &jpeg.Options{jpeg.DefaultQuality})
}
