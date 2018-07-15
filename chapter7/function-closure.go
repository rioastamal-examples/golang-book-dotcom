package main
import "fmt"

func main() {
    var tmpInput int
    var index uint8 = 1

    // Create closure to determine a given number is odd or even
    oddEven := func(number int) string {
        // Closure can access "global" variable where its defined
        // Here we update value of index
        index++

        if number % 2 == 0 {
            return "even"
        }
        return "odd"
    }

    // Put infinite loop
    fmt.Println("Determine number is odd or even.\nType -1 to exit program.")
    fmt.Println("----------------")
    for true {
        fmt.Printf("Try #%d - Input a number: ", index)
        fmt.Scanf("%d", &tmpInput)

        if tmpInput == -1 { break }

        fmt.Printf("Number %v is %v.\n", tmpInput, oddEven(tmpInput))
    }

    fmt.Println("\nSee you later.")
}