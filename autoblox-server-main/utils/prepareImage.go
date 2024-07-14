package utils

import (
	"image/png"
	"io"
	"mime/multipart"

	"github.com/nfnt/resize"
)

func PrepareImage(imgFile *multipart.File, w, h uint) ([]float32, int, int, int, int, error) {
	img, err := png.Decode(*imgFile)

	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	// Resize image
	size := img.Bounds().Size()
	imgWidth, imgHeight := int(size.X), int(size.Y)
	img = resize.Thumbnail(w, h, img, resize.Lanczos3)

	resizedSize := img.Bounds().Size()
	resizedImgWidth, resizedImgHeight := int(resizedSize.X), int(resizedSize.Y)

	red := []float32{}
	green := []float32{}
	blue := []float32{}

	// Get red, green, and blue color for each pixel
	for y := 0; y < int(h); y++ {
		for x := 0; x < int(w); x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			red = append(red, float32(r/257)/255.0)
			green = append(green, float32(g/257)/255.0)
			blue = append(blue, float32(b/257)/255.0)
		}
	}

	// Turn it into (R, G, B)
	rgb := append(red, green...)
	rgb = append(rgb, blue...)

	// Set file cursor back to start
	_, err = (*imgFile).Seek(0, io.SeekStart)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	return rgb, resizedImgWidth, resizedImgHeight, imgWidth, imgHeight, nil
}
