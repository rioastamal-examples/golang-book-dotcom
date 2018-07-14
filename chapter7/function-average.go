package main
import "fmt"

func main() {
    scores := make([]float64, 0)
    var total float64
    var ask_user bool = true
    var index uint8 = 0
    var tmp_input float64

    // Put infinite loop
    fmt.Println("Getting average of score.\nType -1 to quit.")
    fmt.Println("----------------")
    for ask_user {
        fmt.Print("Input score ", index + 1, ": ")

        fmt.Scanf("%f", &tmp_input)

        if (tmp_input == -1) { break }

        // Resize score slice
        scores = append(scores, tmp_input)
        total += tmp_input
        index++
    }

    fmt.Println("-------------")
    fmt.Printf("We have %d score(s) with total %v.\n", index, total)
    fmt.Printf("Average score is %v.\n", avg(scores))
}

// Function start with keyword funcz
// Argument should be defined using: VAR_NAME TYPE, VAR_NAME TYPE, ...
// The type of return value of the function should be defined

/**
 * Function to calculate average of given numbers.
 *
 * @param []float64 numbers
 * @return float64
 */
func avg(numbers []float64) float64 {
    var total float64 = 0.0

    if len(numbers) == 0 {
        // Trigger run time error with panic() function
        panic("Division by zero")
    }

    // We are only interested in value
    for _, value := range numbers {
        total += value
    }

    // Need to convert length of the numbers so it has
    // same data type float64
    return total / float64( len(numbers) )
}