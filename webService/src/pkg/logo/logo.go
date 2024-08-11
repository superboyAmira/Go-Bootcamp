package logo

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func Create() {
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))

	lightBlue := color.RGBA{173, 216, 230, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{lightBlue}, image.Point{}, draw.Src)

	brown := color.RGBA{139, 69, 19, 255}
	for y := 150; y < 250; y++ {
		for x := 100; x < 200; x++ {
			img.Set(x, y, brown)
		}
	}

	darkBrown := color.RGBA{101, 67, 33, 255}
	for y := 200; y < 250; y++ {
		for x := 130; x < 170; x++ {
			img.Set(x, y, darkBrown)
		}
	}

	white := color.RGBA{255, 255, 255, 255}
	for y := 170; y < 190; y++ {
		for x := 110; x < 130; x++ {
			img.Set(x, y, white)
		}
	}
	for y := 170; y < 190; y++ {
		for x := 170; x < 190; x++ {
			img.Set(x, y, white)
		}
	}

	// Сохраняем изображение в файл
	file, err := os.Create("../../web/static/amazing_logo.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)
}