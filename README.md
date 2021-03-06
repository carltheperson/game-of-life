# Game of Life

This is my take on Conway's Game of Life.
It is written in Go using the 2D library [Pixel](https://github.com/faiface/pixel).

![](images/preview-random.gif) ![](images/preview.gif)

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

_Note that the game starts paused with no cells_

# Controls

| Description            | Key           |
| ---------------------- | ------------- |
| Pausing/unpausing      | _SPACE_       |
| Fill with random cells | _R_           |
| Clear all cells        | _C_           |
| Place cell             | _LEFT CLICK_  |
| Remove cell            | _RIGHT CLICK_ |

# Configurations

At the top of the `main.go` file you will find all the configurations.
Change their value to change the functionality of the game.

They currently look like this.

```go
var (
	title           = "Game of life"
	windowResizable = false
	fps             = 10
	cellSize        = 10.0
	cellAmount      = 50
	padding         = 1.0
	cellColor       = colornames.Darkcyan
	backgroundColor = colornames.Lightcyan
)
```

Here is an example with different colors.

![](images/screenshot.png)

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
