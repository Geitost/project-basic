package utils

import (
	"math"
)

func ResizeRgb(input []float32, width, height int) []float32 {
	// Calculate scaling factors
	scaleX := float64(len(input)/3) / float64(width)
	scaleY := 1.0 / float64(height)

	// Output array for the resized image
	output := make([]float32, width*height*3) // Assuming RGB values

	// Iterate over each pixel in the output
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Coordinates in the input array
			srcX := float64(x) * scaleX
			srcY := float64(y) * scaleY

			// Iterate over RGB channels without vectorization
			for c := 0; c < 3; c++ {
				// Calculate Lanczos3 weighted sum for the current channel
				var sum float64
				for i := math.Floor(srcX) - 2; i <= math.Floor(srcX)+2; i++ {
					weightX := lanczos3Interpolation(srcX - i)
					for j := math.Floor(srcY) - 2; j <= math.Floor(srcY)+2; j++ {
						weightY := weightX * lanczos3Interpolation(srcY-j)
						idx := int(j)*width*3 + int(i)*3 + c
						if idx >= 0 && idx+2 < len(input) {
							sum += float64(input[idx]) * weightY
						}
					}
				}
				output[y*width*3+x*3+c] = float32(sum)
			}
		}
	}

	return output
}

func lanczos3Interpolation(x float64) float64 {
	if x == 0 {
		return 1
	}
	if math.Abs(x) < 3 {
		return math.Sin(math.Pi*x) * math.Sin(math.Pi*x/3) / (math.Pi * math.Pi * x * x / 3)
	}
	return 0
}
