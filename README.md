# Game of Life

This is my take on Conway's Game of Life.
It is written in Go using the 2D library [Pixel](https://github.com/faiface/pixel).

![](preview.gif)

# Running

You need OpenGL development libraries to run this. Read more [here](https://github.com/faiface/pixel#requirements).

Clone the repository.

```bash
git clone https://github.com/carltheperson/game-of-life
```

Install dependencies.

```bash
go get
```

Run project.

```bash
go run .
```

# Configurations

At the top of the `main.go` file you will find all the configurations.
Change their value to change to functionality of the game.

They currently look like this.

```go
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
```

Here is an example with different colors.

![](screenshot.png)

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
