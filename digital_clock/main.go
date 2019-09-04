package main

import (
    "fmt"
    "time"
    "strconv"
)

func main() {
    type placeholder [5]string

    zero := placeholder{
        "███",
        "█ █",
        "█ █",
        "█ █",
        "███",
    }
    one := placeholder{
        "██ ",
        " █ ",
        " █ ",
        " █ ",
        "███",
    }
    two := placeholder{
        "███",
        "  █",
        "███",
        "█  ",
        "███",
    }
    three := placeholder{
        "███",
        "  █",
        "███",
        "  █",
        "███",
    }
    four := placeholder{
        "█ █",
        "█ █",
        "███",
        "  █",
        "  █",
    }
    five := placeholder{
        "███",
        "█  ",
        "███",
        "  █",
        "███",
    }
    six := placeholder{
        "███",
        "█  ",
        "███",
        "█ █",
        "███",
    }
    seven := placeholder{
        "███",
        "  █",
        "  █",
        "  █",
        "  █",
    }
    eight := placeholder{
        "███",
        "█ █",
        "███",
        "█ █",
        "███",
    }
    nine := placeholder{
        "███",
        "█ █",
        "███",
        "  █",
        "███",
    }
    seperator := [...]placeholder{
        {
            "   ",
            " ░ ",
            "   ",
            " ░ ",
            "   ",
        },
        {
            "   ",
            "   ",
            "   ",
            "   ",
            "   ",
        },
    }

    digits := [...]placeholder{
        zero,
        one,
        two,
        three,
        four,
        five,
        six,
        seven,
        eight,
        nine,
    }

    var (
        hr int
        min int
        sec int
        output string
        digit int
        toggle int
    )

    const (
        clear = "\033[H\033[2J"
        COLON = 0
        VOID  = 1
    )

    for {
        hr  = time.Now().Hour()
        min = time.Now().Minute()
        sec = time.Now().Second()

        output = fmt.Sprintf("%02d:%02d:%02d", hr, min, sec)

        fmt.Print(clear)

        for i := 0; i < 5; i++ {
            for _, v := range output {
                if string(v) == ":" {
                    fmt.Printf("%s ", seperator[toggle][i])
                    continue
                }
                digit, _ = strconv.Atoi(string(v))
                fmt.Printf("%s ", digits[digit][i])
            }
            fmt.Println()
        }

        toggle++
        if toggle == 2 {
            toggle = 0
        }
        time.Sleep(time.Second)
    }
}
