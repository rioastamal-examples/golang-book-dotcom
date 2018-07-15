package main
import "fmt"

// Slice is an array which the length can be dynamically changed
func main() {
    // Create an empty array with 8 bits unsigned integer
    var numbers []uint8
    fmt.Println(numbers)

    // Create slice using make() function
    numbers = make([]uint8, 5)
    fmt.Println(numbers)

    // Resize the length to 3
    numbers = make([]uint8, 3)
    fmt.Println(numbers)
    fmt.Println("Length of numbers is", len(numbers))

    scores := make([]float64, 0)
    var total float64
    var askUser bool = true
    var index uint8 = 0
    var tmpInput float64

    // Put infinite loop
    fmt.Println()
    for askUser {
        fmt.Print("Input score ", index + 1, ": ")

        fmt.Scanf("%f", &tmpInput)

        if (tmpInput == -1) {
            askUser = false
        } else {
            // Resize score slice
            scores = append(scores, tmpInput)
            total += tmpInput
            index++
        }
    }

    fmt.Println(scores)
    fmt.Println("Average score =", total / float64(index))
}