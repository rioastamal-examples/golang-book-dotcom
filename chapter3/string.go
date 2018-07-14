package main
import "fmt"

func main() {
    var x string
    // Direct assignment
    var y string = "World"
    // Auto detection of variable type
    z := "For All"
    x = "Hello "
    fmt.Println(x + y + " " + z)
}