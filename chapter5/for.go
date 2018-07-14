package main
import "fmt"

func main() {
    // Small integer 8 bits = 1 byte
    var i uint8 = 1

    // For loop type 1
    fmt.Println("for i <= 10 {...}")
    for i <= 10 {
        fmt.Println("Loop", i)
        i += 1
    }

    // For loop type 2
    fmt.Println("\nfor i = 1; i <= 10; i++ {...}")
    for i = 1; i <= 10; i++ {
        fmt.Println("Loop", i)
    }
}