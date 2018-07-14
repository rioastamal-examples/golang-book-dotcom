package main
import "fmt"

func main() {
    // We support input for decimal and floating
    var (
        input float64
        output float64
    )

    fmt.Print("Enter a number: ")

    // %f for Floating number and assign to input variable
    fmt.Scanf("%f", &input)

    output = input * 2
    fmt.Println(input, "x 2 = ", output)
}