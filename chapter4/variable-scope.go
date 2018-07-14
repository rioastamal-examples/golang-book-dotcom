package main
import "fmt"

// Variable scope are exists between { } and it tested { }

// x below is global variable in a package
var x string = "Hello World"

func main() {
    {
        y := "I am private"
        fmt.Println(y)
    }
    fmt.Println(x)

    // uncomment line below to see the error
    // fmt.Println(y)
}