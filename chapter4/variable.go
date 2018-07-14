package main
import "fmt"

func main() {
    var x string
    // Direct assignment
    var y string = "World"
    // Auto detection of variable type. No need var keyword.
    z := "For All"
    x = "Hello "
    fmt.Println(x + y + " " + z)
}