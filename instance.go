package image

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

type Image struct {
	imagePath string
}

func New(imagePath string) *Image {
	return &Image{imagePath: imagePath}
}

func (img *Image) wrapRotate(rotate func(targetImage image.Image) image.Image, imagePath ...string) error {
	file, err := os.Open(img.imagePath)
	if err != nil {
		return fmt.Errorf("open the file error: %v", err)
	}
	defer file.Close()

	targetImage, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("decode the image file failed: %v", err)
	}

	newImage := rotate(targetImage)

	var writePath string
	if len(imagePath) > 0 {
		writePath = imagePath[0]
	} else {
		writePath = img.imagePath
	}
	file, err = os.Create(writePath)
	if err != nil {
		return fmt.Errorf("open the image path failed: %v", err)
	}
	defer file.Close()

	// write image
	if err := png.Encode(file, newImage); err != nil {
		return fmt.Errorf("encode the image into file failed: %v", err)
	}

	return nil
}

func (img *Image) CWRotate90(imagePath ...string) error {
	// Rotate 90 degrees clockwise
	rotate := func(targetImage image.Image) image.Image {
		newImage := image.NewRGBA(image.Rect(0, 0, targetImage.Bounds().Dy(), targetImage.Bounds().Dx()))
		for y := targetImage.Bounds().Min.Y; y < targetImage.Bounds().Max.Y; y++ {
			for x := targetImage.Bounds().Min.X; x < targetImage.Bounds().Max.X; x++ {
				newImage.Set(targetImage.Bounds().Max.Y-y, x, targetImage.At(x, y))
			}
		}
		return newImage
	}
	return img.wrapRotate(rotate, imagePath...)
}

func (img *Image) CWRotate180(imagePath ...string) error {
	rotate := func(targetImage image.Image) image.Image {
		newImage := image.NewRGBA(image.Rect(0, 0, targetImage.Bounds().Dx(), targetImage.Bounds().Dy()))
		for y := targetImage.Bounds().Min.Y; y < targetImage.Bounds().Max.Y; y++ {
			for x := targetImage.Bounds().Min.X; x < targetImage.Bounds().Max.X; x++ {
				newImage.Set(targetImage.Bounds().Max.X-x, targetImage.Bounds().Max.Y-y, targetImage.At(x, y))
			}
		}
		return newImage
	}
	return img.wrapRotate(rotate, imagePath...)
}

func (img *Image) CWRotate270(imagePath ...string) error {
	rotate := func(targetImage image.Image) image.Image {
		newImage := image.NewRGBA(image.Rect(0, 0, targetImage.Bounds().Dy(), targetImage.Bounds().Dx()))
		for y := targetImage.Bounds().Min.Y; y < targetImage.Bounds().Max.Y; y++ {
			for x := targetImage.Bounds().Min.X; x < targetImage.Bounds().Max.X; x++ {
				newImage.Set(y, targetImage.Bounds().Max.X-x, targetImage.At(x, y))
			}
		}
		return newImage
	}
	return img.wrapRotate(rotate, imagePath...)
}
