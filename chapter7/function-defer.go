package main
import "fmt"

func main() {
    var tmpInput [2]int

    // Calls goodBye() function using defer. So it is automatically
    // gets called when the function execution ends. Even if panic() occurs.
    // It similar to Destruct() method on Object Oriented.
    // In real world app, defer mostly used to close or cleanup I/O resource
    // which has been opened.
    defer goodBye()

    fmt.Println("Division of two numbers.")
    fmt.Print("Enter first number: ")
    fmt.Scanf("%d", &tmpInput[0])
    fmt.Print("Enter second number (divider): ")
    fmt.Scanf("%d", &tmpInput[1])

    fmt.Printf("Division result %v.\n", tmpInput[0] / tmpInput[1])
}

func goodBye() {
    fmt.Println("---- See you later. Bye! ----")
}