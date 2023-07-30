# goimage
the image to rotate by golang, including rotate 90, 180 and 270 degress clockwise.

# usage

``` golang

import (
    "github.com/quixote-liu/goimage"
)

func main() {
    imagePath := "./test/test.png"
    outputPath := "./test/test_90.png"
    err := goimage.CWRotate90(imagePath, outputPath)
    if err != nil {
        panic(err)
    }
}

```
