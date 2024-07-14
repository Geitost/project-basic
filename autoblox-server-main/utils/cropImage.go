package utils

import (
	"image"
	"image/png"
	"io"
	"mime/multipart"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

func CropAndPrepareImage(imgFile *multipart.File, x, y, cropW, cropH, w, h uint) ([]float32, int64, int64, int64, int64, error) {
	img, err := png.Decode(*imgFile)

	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	// Resize image
	size := img.Bounds().Size()
	imgWidth, imgHeight := int64(size.X), int64(size.Y)
	img, err = cutter.Crop(img, cutter.Config{
		Width:  int(cropW),
		Height: int(cropH),
		Anchor: image.Point{int(x), int(y)},
	})
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	// Height is set to 0 to keep aspect ratio
	img = resize.Resize(w, h, img, resize.Lanczos3)

	resizedSize := img.Bounds().Size()
	resizedImgWidth, resizedImgHeight := int64(resizedSize.X), int64(resizedSize.Y)

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
