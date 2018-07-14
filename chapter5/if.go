package main
import "fmt"

func main() {
    var name string

    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    if name == "Rio" {
        fmt.Println("Welcome Rio!")
    } else if name == "Wonokairun" {
        fmt.Println("Long time no see Wono!")
    } else {
        fmt.Println("I'm sorry, do I know you?")
    }
}