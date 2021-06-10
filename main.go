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
	cellAmount      = 50
	padding         = 1.0
	cellColor       = pixel.RGB(0, 0, 0)
	backgroundColor = colornames.Skyblue
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
	if isAliveNow && numberOfNeighbors == 2 || isAliveNow && numberOfNeighbors == 3 {
		return true
	}

	if !isAliveNow && numberOfNeighbors == 3 {
		return true
	}

	if isAliveNow {
		return false
	}

	return isAliveNow
}

func generateCells() [][]*imdraw.IMDraw {
	x := make([][]*imdraw.IMDraw, cellAmount)
	for i := 0; i < cellAmount; i++ {
		y := make([]*imdraw.IMDraw, cellAmount)
		for j := 0; j < cellAmount; j++ {
			coordX := float64(i)*cellSize + float64(i)*padding
			coordY := float64(j)*cellSize + float64(j)*padding
			imd := imdraw.New(nil)
			imd.Color = cellColor
			imd.Push(pixel.V(coordX, coordY))
			imd.Push(pixel.V(coordX+cellSize, coordY))
			imd.Push(pixel.V(coordX+cellSize, coordY+cellSize))
			imd.Push(pixel.V(coordX, coordY+cellSize))
			imd.Polygon(0)
			y[j] = imd
		}
		x[i] = y

	}
	return x
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

	cells := generateCells()
	matrix := generateRandomMatrix()

	start := time.Now()
	for !win.Closed() {
		win.Clear(backgroundColor)

		elapsed := float64(time.Since(start).Milliseconds())
		if elapsed > timeForOneFrameMilliseconds {
			start = time.Now()
			matrix = updateMatrix(matrix)
		}

		for i, yBlocks := range cells {
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
	pixelgl.Run(run)
}
