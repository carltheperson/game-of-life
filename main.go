package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	title           = "Game of life"
	windowResizable = false
	fps             = 10
	cellSize        = 10.0
	cellAmount      = 80
	padding         = 1.0
	cellColor       = colornames.Darkcyan
	backgroundColor = colornames.Lightcyan
)

var windowSize = cellSize*float64(cellAmount) + padding*float64(cellAmount)
var cfg = pixelgl.WindowConfig{
	Title:     title,
	Resizable: windowResizable,
	Bounds:    pixel.R(0, 0, float64(windowSize), float64(windowSize)),
	VSync:     true,
}

var neighborCoords = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

var timeForOneFrameMilliseconds = (1.0 / float64(fps)) * 1000

func getIsCellAliveNextRound(isAliveNow bool, numberOfNeighbors int) bool {
	// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	if isAliveNow && numberOfNeighbors < 2 {
		return false
	}

	// Any live cell with two or three live neighbours lives on to the next generation.
	if isAliveNow && (numberOfNeighbors == 2 || numberOfNeighbors == 3) {
		return true
	}

	// Any live cell with more than three live neighbours dies, as if by overpopulation.
	if isAliveNow && numberOfNeighbors > 3 {
		return false
	}

	// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	if !isAliveNow && numberOfNeighbors == 3 {
		return true
	}

	return isAliveNow
}

func drawCells(matrix [][]bool, win *pixelgl.Window) {
	imd := imdraw.New(nil)
	for i := 0; i < cellAmount; i++ {
		for j := 0; j < cellAmount; j++ {
			if !matrix[i][j] {
				continue
			}
			coordX := float64(i)*cellSize + float64(i)*padding
			coordY := float64(j)*cellSize + float64(j)*padding
			imd.Color = cellColor
			imd.Push(pixel.V(coordX, coordY))
			imd.Push(pixel.V(coordX+cellSize, coordY+cellSize))
			imd.Rectangle(0)
		}

	}
	imd.Draw(win)
}

func generateMatrix() [][]bool {
	x := make([][]bool, cellAmount)
	for i := 0; i < cellAmount; i++ {
		y := make([]bool, cellAmount)
		for j := 0; j < cellAmount; j++ {
			y[j] = false
		}
		x[i] = y
	}
	return x
}

func generateRandomMatrix() [][]bool {
	rand.Seed(time.Now().Unix())
	matrix := generateMatrix()

	for i := range matrix {
		for j := range matrix[0] {
			if rand.Intn(2) == 1 {
				matrix[i][j] = true
			} else {
				matrix[i][j] = false
			}
		}
	}
	return matrix
}

func checkNeighbor(i int, j int, matrix [][]bool) bool {
	if i >= 0 && j >= 0 && i < len(matrix) && j < len(matrix[0]) {
		return matrix[i][j]
	}
	return false
}

func getNumberOfNeighbors(i int, j int, matrix [][]bool) int {
	numberOfNeighbors := 0
	for _, coords := range neighborCoords {
		if checkNeighbor(i+coords[0], j+coords[1], matrix) {
			numberOfNeighbors += 1
		}
	}
	return numberOfNeighbors
}

func updateMatrix(matrix [][]bool) [][]bool {
	newMatrix := generateMatrix()
	for i := range matrix {
		for j := range matrix[0] {
			numberOfNeighbors := getNumberOfNeighbors(i, j, matrix)
			newMatrix[i][j] = getIsCellAliveNextRound(matrix[i][j], numberOfNeighbors)

		}
	}
	return newMatrix
}

func run() {
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	matrix := generateMatrix()

	paused := true

	start := time.Now()
	for !win.Closed() {
		win.Clear(backgroundColor)

		elapsed := float64(time.Since(start).Milliseconds())
		if elapsed > timeForOneFrameMilliseconds && !paused {
			start = time.Now()
			matrix = updateMatrix(matrix)
		}

		drawCells(matrix, win)
		if win.Pressed(pixelgl.MouseButton1) || win.Pressed(pixelgl.MouseButton2) {
			x := win.MousePosition().X
			y := win.MousePosition().Y
			i := int((x / windowSize) * float64(cellAmount))
			j := int((y / windowSize) * float64(cellAmount))
			if win.Pressed(pixelgl.MouseButton1) {
				matrix[i][j] = true
			} else {
				matrix[i][j] = false
			}
		}

		if win.JustPressed(pixelgl.KeySpace) {
			paused = !paused
		}

		if win.JustPressed((pixelgl.KeyR)) {
			matrix = generateRandomMatrix()
		}

		if win.JustPressed((pixelgl.KeyC)) {
			matrix = generateMatrix()
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
