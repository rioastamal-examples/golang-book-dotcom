package main
import "fmt"

func main() {
    var name string

    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    switch name {
        case "Rio":
            fmt.Println("Welcome Rio!")
        case "Wonokairun":
            fmt.Println("Long time no see Wono!")
        default:
            fmt.Println("I'm sorry, do I know you?")
    }
}