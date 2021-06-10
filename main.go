package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func generateBlocks(w int, l int) [][]*imdraw.IMDraw {
	blockSize := 10.0
	padding := 2.0
	x := make([][]*imdraw.IMDraw, w)
	for i := 0; i < w; i++ {
		y := make([]*imdraw.IMDraw, l)
		for j := 0; j < l; j++ {
			coordX := float64(i)*blockSize + float64(i)*padding
			coordY := float64(j)*blockSize + float64(j)*padding
			imd := imdraw.New(nil)
			imd.Circle(coordX, coordY)
			imd.Color = pixel.RGB(1, 0, 0)
			imd.Ellipse(pixel.V(120, 80), 0)
			imd.Push(pixel.V(coordX, coordY))
			imd.Push(pixel.V(coordX+blockSize, coordY))
			imd.Push(pixel.V(coordX+blockSize, coordY+blockSize))
			imd.Push(pixel.V(coordX, coordY+blockSize))
			imd.Polygon(0)
			y[j] = imd
		}
		x[i] = y

	}
	return x
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game of life",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	blocks := generateBlocks(5, 6)

	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		for _, yBlocks := range blocks {
			for _, block := range yBlocks {
				block.Draw(win)
			}
		}
		win.Update()
	}
}

func main() {
	generateBlocks(6, 5)
	pixelgl.Run(run)
}
