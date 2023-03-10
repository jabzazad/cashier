package resize

import (
	"image"

	"github.com/disintegration/imaging"
)

// Image resize image
func Image(originalImage image.Image, maxWidth int) *image.NRGBA {
	if originalImage.Bounds().Size().X > maxWidth {
		return imaging.Resize(originalImage, maxWidth, 0, imaging.Lanczos)
	}
	return imaging.Resize(originalImage, originalImage.Bounds().Size().X, 0, imaging.Lanczos)
}
