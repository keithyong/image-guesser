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
    
    newImg, _ := os.Create("new.jpg")
    defer newImg.Close()
    
    b := img.Bounds()
    m := image.NewRGBA(image.Rect(0, 0, b.Max.X, b.Max.Y * 2))
    
    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {
           m.Set(x, y, img.At(x, y))
           m.Set(b.Max.X - x, y, img.At(x, y))
        }
        fmt.Println(y)
    }
    
    jpeg.Encode(newImg, m, &jpeg.Options{jpeg.DefaultQuality})
}
