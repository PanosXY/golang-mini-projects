package main

import (
    "fmt"
    "time"
    "math/rand"
    "os"
    "strings"
    "strconv"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    if len(os.Args) < 3 || os.Args[2] != "positive" && os.Args[2] != "negative" {
        fmt.Println("[name] [positive|negative]")
        return
    }

    moods := [...][3]string{
        {":)", ":D",  "xD"},
        {":(", ":-(", ":["},
    }

    name := os.Args[1]
    moodType := os.Args[2]
    var idx int
    switch moodType {
    case "positive":
        idx = 0
    case "negative":
        idx = 1
    }
    mood := moods[idx][rand.Intn(len(moods[idx]))]
    fmt.Printf("%s feels %s\n", name, mood)


    var (
        names     [3]string
        distances [5]int
        data      [5]byte
        ratios    [1]float64
        alives    [4]bool
        zero      [0]byte
    )

    fmt.Printf("names    : %#v\n", names)
    fmt.Printf("distances: %#v\n", distances)
    fmt.Printf("data     : %#v\n", data)
    fmt.Printf("rations  : %#v\n", ratios)
    fmt.Printf("alives   : %#v\n", alives)
    fmt.Printf("zero     : %#v\n", zero)
    fmt.Println()
    fmt.Printf("names    : %T\n", names)
    fmt.Printf("distances: %T\n", distances)
    fmt.Printf("data     : %T\n", data)
    fmt.Printf("rations  : %T\n", ratios)
    fmt.Printf("alives   : %T\n", alives)
    fmt.Printf("zero     : %T\n", zero)
    fmt.Println()
    fmt.Printf("names    : %q\n", names)
    fmt.Printf("distances: %d\n", distances)
    fmt.Printf("data     : %d\n", data)
    fmt.Printf("rations  : %.2f\n", ratios)
    fmt.Printf("alives   : %t\n", alives)
    fmt.Printf("zero     : %d\n", zero)


    nnames := [3]string{"Einstein", "Shepard", "Tesla"}
    books := [5]string{"Kafka's Revenge", "Stay Golden"}

    fmt.Printf("%q\n", nnames)
    fmt.Printf("%q\n", books)


    week := [4]string{"Monday", "Tuesday"}
    wend := [4]string{"Saturday", "Sunday"}

    fmt.Println(week != wend)

    evens := [...]int{2, 4, 6, 8, 10}
    evens2 := [...]int{2, 4, 6, 8, 10}

    fmt.Println(evens == evens2)

    image := [5]byte{'h', 'i'}
    var ddata [5]uint8

    fmt.Println(ddata == image)


    bbooks := [...]string{"xyZ", "aBc", "KLm"}
    upper := bbooks
    lower := bbooks

    for i, v := range bbooks {
        upper[i] = strings.ToUpper(v)
        lower[i] = strings.ToLower(v)
    }

    fmt.Printf("books: %#v\n", bbooks)
    fmt.Printf("upper: %#v\n", upper)
    fmt.Printf("lower: %#v\n", lower)


    scientists := [...][3]string{
        {"Albert",  "Einstein", "time"},
        {"Isaac",   "Newton",   "apple"},
        {"Stephen", "Hawking",  "blackhole"},
        {"Marie",   "Curie",    "radium"},
        {"Charles", "Darwin",   "fittest"},
    }

    fmt.Println("First Name\tLast Name\tNickname")
    fmt.Println("==========================================")
    for _, scientist := range scientists {
        for _, v := range scientist {
            fmt.Printf("%s\t\t", v)
        }
        fmt.Println()
    }

    if len(os.Args) < 4 {
        fmt.Println("Please provide the amount to be converted")
        return
    }

    amount, err := strconv.ParseFloat(os.Args[3], 64)
    if err != nil {
        fmt.Println("It should be a number")
        return
    }

    const (
        EUR = iota
        GBP
        JPY
    )

    convRatios := [...]float64{
        EUR: 0.9,
        JPY: 106.16,
        GBP: 0.82,
    }

    fmt.Printf("%g USD is %.2f EUR\n", amount, amount * convRatios[EUR])
    fmt.Printf("%g USD is %.2f GBP\n", amount, amount * convRatios[GBP])
    fmt.Printf("%g USD is %.2f JPY\n", amount, amount * convRatios[JPY])


    if len(os.Args) < 5 {
        fmt.Println("Tell me a book title")
        return
    }

    bookshelf := [...]string{
        "Kafka's Revenge",
        "Stay Golden",
        "Everythingship",
        "Kafka's Revenge 2nd Edition",
    }

    sq := strings.ToLower(os.Args[4])
    found := false
    fmt.Println("Search Results:")
    for _, v := range bookshelf {
        if strings.Contains(strings.ToLower(v), sq) {
            fmt.Printf("+ %s\n", v)
            found = true
        }
    }
    if !found {
        fmt.Printf("We don't have book: %s\n", os.Args[4])
    }


    numbers := os.Args[5:]
    if len(numbers) < 1 || len(numbers) > 5 {
        fmt.Println("Maximum numbers: 5, Minimum: 1")
        return
    }

    var (
        arrayTBF [5]float64
        avg float64
        validNums int
    )
    for i, v := range numbers {
        if n, err := strconv.ParseFloat(v, 64); err != nil {
            continue
        } else {
            arrayTBF[i] = n
            avg += n
            validNums++
        }
    }
    avg /= float64(validNums)
    fmt.Printf("Your numbers: %g\n", arrayTBF)
    fmt.Printf("Average: %.2f\n", avg)

    for range arrayTBF {
        for j := range arrayTBF {
            if j == len(arrayTBF) - 1 {
                break
            }
            if arrayTBF[j] > arrayTBF[j+1] {
                arrayTBF[j], arrayTBF[j+1] = arrayTBF[j+1], arrayTBF[j]
            }
        }
    }
    fmt.Printf("The sorted array is: %g\n", arrayTBF)
}
