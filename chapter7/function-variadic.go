package main
import "fmt"

func main() {
    sum, len, numbers := sumNumbers(10, 5, 10, 5, 10)

    fmt.Printf("Sum of the numbers: %v.\n", sum)
    fmt.Printf("Total number given: %v.\n", len)
    fmt.Println(numbers)
}

/**
 * Function to sum given numbers.
 *
 * @param ...int numbers
 * @return int Sum of the numbers
 * @return int How many numbers given
 */
func sumNumbers(numbers ...int) (theSum int, theLen int, theNumbers []int) {
    for _, value := range numbers {
        theSum += value
    }
    theLen = len(numbers)
    theNumbers = numbers

    return
}