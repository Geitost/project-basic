package utils

func CropRgb(input []float32, x1, y1, x2, y2, originalWidth, originalHeight int) []float32 {
	// Calculate width and height from coordinates
	width := x2 - x1
	height := y2 - y1

	// Output array for the cropped image
	output := make([]float32, width*height*3) // Assuming RGB values

	// Iterate over each pixel in the output
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			// Coordinates in the input array
			srcX := x1 + i
			srcY := y1 + j

			// Copy RGB values without vectorization
			srcIdx := (srcY*originalWidth + srcX) * 3
			dstIdx := (j*width + i) * 3

			// Ensure indices are within bounds
			if srcIdx >= 0 && srcIdx+2 < len(input)*3 && dstIdx >= 0 && dstIdx+2 < len(output) {
				// Copy RGB values
				output[dstIdx] = input[srcIdx]
				output[dstIdx+1] = input[srcIdx+1]
				output[dstIdx+2] = input[srcIdx+2]
			}
		}
	}

	return output
}
