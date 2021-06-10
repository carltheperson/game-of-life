package main

import (
	"math/rand"
	"time"

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
			imd.Color = pixel.RGB(0, 0, 0)
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

func generateMatrix(w int, l int) [][]bool {
	x := make([][]bool, w)
	for i := 0; i < w; i++ {
		y := make([]bool, l)
		x[i] = y
	}
	return x
}

func generateRandomMatrix(w int, l int) [][]bool {
	rand.Seed(time.Now().Unix())
	x := make([][]bool, w)
	for i := 0; i < w; i++ {
		y := make([]bool, l)
		for j := 0; j < l; j++ {
			if rand.Intn(2) == 1 {
				y[j] = true
			} else {
				y[j] = false
			}
		}
		x[i] = y
	}
	return x
}

func checkNeighbor(i int, j int, matrix [][]bool) bool {
	if i >= 0 && j >= 0 && i < len(matrix) && j < len(matrix[0]) {
		return matrix[i][j]
	}
	return false
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

	blocks := generateBlocks(50, 60)
	matrix := generateRandomMatrix(50, 60)

	neighborCoords := [][]int{{-1, -1}, {0, -1}, {1, -1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	fps := 10
	timeForOneFrameMilliseconds := (1.0 / float64(fps)) * 1000

	start := time.Now()
	for !win.Closed() {
		win.Clear(colornames.Skyblue)

		t := time.Now()
		elapsed := float64(t.Sub(start).Milliseconds())
		if elapsed > timeForOneFrameMilliseconds {
			start = time.Now()
			newMatrix := generateMatrix(50, 60)
			for i := range matrix {
				for j := range matrix[0] {
					numberOfNeighbors := 0
					for _, coords := range neighborCoords {
						if checkNeighbor(i+coords[0], j+coords[1], matrix) {
							numberOfNeighbors += 1
						}
					}

					if matrix[i][j] && numberOfNeighbors < 2 {
						newMatrix[i][j] = false
					} else if !matrix[i][j] && numberOfNeighbors == 3 {
						newMatrix[i][j] = true
					} else if matrix[i][j] && numberOfNeighbors > 3 {
						newMatrix[i][j] = false
					} else {
						newMatrix[i][j] = matrix[i][j]
					}

				}
			}
			matrix = newMatrix
		}

		for i, yBlocks := range blocks {
			for j, block := range yBlocks {
				if matrix[i][j] {
					block.Draw(win)
				}
			}
		}

		win.Update()
	}
}

func main() {
	generateBlocks(6, 5)
	pixelgl.Run(run)
}
