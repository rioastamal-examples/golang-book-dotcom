package main
import "fmt"

func main() {
    // Static integer arrays. Default value is 0 for each element
    var numbers [5]uint8
    numbers[0] = 11
    numbers[1] = 22
    fmt.Println(numbers)

    // Static string arrays -- shortcut
    names := [3]string{ "Wonokairun", "Bunali", "Brudin" }
    fmt.Println(names, "\n")

    var scores [3]float64
    var total float64
    var i uint8

    for i = 0; i < 3; i++ {
        fmt.Print("Input score ", i + 1, ": ")
        fmt.Scanf("%f", &scores[i])
    }

    fmt.Println()
    for i = 0; i < 3; i++ {
        fmt.Println("Score", i + 1, "=", scores[i])
        total += scores[i]
    }

    // Convert len of scores to float64
    fmt.Println("Average score =", total / float64(len(scores)))
}