package main
import "fmt"

func main() {
    scores := make([]float64, 0)
    var index uint8 = 0
    var tmp_input float64

    // Put infinite loop
    fmt.Println("Getting average of score.\nType -1 to quit.")
    fmt.Println("----------------")
    for true {
        fmt.Print("Input score ", index + 1, ": ")

        fmt.Scanf("%f", &tmp_input)

        if tmp_input == -1 { break }

        // Resize score slice
        scores = append(scores, tmp_input)
        index++
    }

    avg_result, total, odd_even, avg_result_int := avg(scores)
    fmt.Println("-------------")
    fmt.Printf("We have %d score(s) with total %v.\n", index, total)
    fmt.Printf("Average score is %v.\n", avg_result)
    fmt.Printf("Average score is %v number (integer -> %v).\n", odd_even, avg_result_int)
}

// Function start with keyword funcz
// Argument should be defined using: VAR_NAME TYPE, VAR_NAME TYPE, ...
// The type of return value of the function should be defined

/**
 * Function to calculate average of given numbers.
 *
 * @param []float64 numbers
 * @return float64 average score
 * @return float64 sum score
 * @return string odd or even number
 * @return int convert the avg to int
 */
func avg(numbers []float64) (float64, float64, string, int) {
    var total float64
    var result float64

    if len(numbers) == 0 {
        // Trigger run time error with panic() function
        panic("Division by zero")
    }

    // We are only interested in value
    for _, value := range numbers {
        total += value
    }

    // Return values
    // 1. Average score. Need to convert length of the numbers so it has
    //    same data type float64
    // 2. Sum of the score
    // 3. Even or odd number
    // 4. Integer version of the result
    result = total / float64( len(numbers) )

    if int(result) % 2 == 0 {
        return result, total, "even", int(result)
    }
    return result, total, "odd", int(result)
}