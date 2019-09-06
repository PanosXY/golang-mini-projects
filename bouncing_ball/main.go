package main

import (
    "fmt"
    "time"
    "os"
    "os/exec"
    "strings"
    "strconv"
)

func main() {
    const (
        ball = 0x1F534
        void = ' '
        clear = "\033[H\033[2J"
        backward = false
        forward = true
        speed = time.Second/20
    )

    var (
        x int
        y int
        prevX int = 1
        prevY int = 1
        xAxisDirection bool
        yAxisDirection bool
        found bool
    )

    // Get terminals' size
    cmd := exec.Command("stty", "size")
    cmd.Stdin = os.Stdin
    out, err := cmd.Output()
    if err != nil {
        fmt.Println("Couldn't take the terminal's size :(")
        return
    }
    size := strings.Fields(string(out))
    height, _ := strconv.Atoi(size[0])
    width, _ := strconv.Atoi(size[1])
    height--
    width--

    // Create board
    board := make([][]bool, height)
    for row := range board {
        board[row] = make([]bool, width)
    }
    buffer := make([]rune, 0, width)

    for {
        fmt.Print(clear)

        // If ball reach board's limit change direction
        if x == height-1 && xAxisDirection == forward {
            xAxisDirection = backward
        }
        if x == 0 && xAxisDirection == backward {
            xAxisDirection = forward
        }

        if y == width-1 && yAxisDirection == forward {
            yAxisDirection = backward
        }
        if y == 0 && yAxisDirection == backward {
            yAxisDirection = forward
        }

        board[x][y] = true
        board[prevX][prevY] = false
        prevX = x
        prevY = y

        if xAxisDirection == backward {
            x--
        } else {
            x++
        }

        if yAxisDirection == backward {
            y--
        } else {
            y++
        }

        found = false
        for i, b := range board {
            for j := range b {
                if board[i][j] {
                    buffer = append(buffer, ball)
                    found = true
                } else {
                    buffer = append(buffer, void)
                }
            }
            fmt.Println(string(buffer))
            buffer = buffer[:0] // Clear buffer
            if found { // For efficiency
                break
            }
        }

        time.Sleep(speed)
    }
}
