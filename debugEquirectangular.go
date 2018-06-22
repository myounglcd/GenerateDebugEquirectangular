package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	const width, height = 512, 256

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var u = float64(x) / float64(width)
			var v = float64(y) / float64(height)

			var long = u * 2 * math.Pi
			var lat = v * math.Pi

			var xx = math.Cos(long) * math.Sin(lat)
			var yy = math.Sin(long) * math.Sin(lat)
			var zz = math.Cos(lat)

			var xxx = math.Abs(xx)
			var yyy = math.Abs(yy)
			var zzz = math.Abs(zz)

			var red = 0
			var green = 0
			var blue = 0

			if (xxx >= yyy) && (xxx >= zzz) {
				red = 255
			}

			if yyy > xxx && yyy >= zzz {
				green = 255
			}

			if zzz > yyy && zzz > xxx {
				blue = 255
			}

			img.Set(x, y, color.NRGBA{
				R: uint8(red),
				G: uint8(green),
				B: uint8(blue),
				A: 255,
			})
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
