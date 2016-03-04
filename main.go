package main;

import (
    "image"
    "image/jpeg"
	"os"
    "errors"
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
    
    m, _ := crop(image.Rect(10, 10, 250, 250), img)
    
    jpeg.Encode(newImg, m, &jpeg.Options{jpeg.DefaultQuality})
}

// Returns a new cropped image
func crop(r image.Rectangle, i image.Image) (*image.RGBA, error) {
    b := i.Bounds()
    var c *image.RGBA
    
    // Out of bounds error checking
    if (b.Max.X < r.Max.X || b.Max.Y < r.Max.Y) {
        err := errors.New("Input rectangle larger than input image bounds.")
        return c, err
    } else if (r.Min.X < b.Min.X || r.Min.Y < b.Min.Y) {
        err := errors.New("Input rectangle smaller than input image bounds.")
        return c, err
    }
    
    c = image.NewRGBA(r)
    
    for y := r.Min.Y; y < r.Max.Y; y++ {
        for x := r.Min.X; x < r.Max.X; x++ {
            c.Set(x, y, i.At(x, y))
        }
    }

    return c, nil
}