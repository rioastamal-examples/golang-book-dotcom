package main
import "fmt"

// Slice is an array which the length can be dynamically changed
func main() {
    // Create an empty map with 8 bits unsigned integer for key
    // and string as its value
    var numbers map[uint8]string

    // Below should be error
    // --> panic: assignment to entry in nil map
    // Because map should initialized with make() first before it can be used
    // numbers[0] = "Zero"
    numbers = make(map[uint8]string)
    numbers[0] = "Zero"
    numbers[1] = "One"
    fmt.Println(numbers)

    // Or like usual we use shortcut assignment
    numbers2 := make(map[uint8]string)
    numbers2[0] = "Zero"
    numbers2[1] = "One"
    fmt.Println(numbers2)

    // Or even shorter
    numbers3 := map[uint8]string{
        0: "Zero",
        1: "One",
        2: "Two", // the last comma is required for multi-line code
    }
    fmt.Println(numbers3)

    // Create users map which contain another map
    // consist of first_name and last_name
    users := make(map[string]map[string]string)
    fmt.Println()

    var tmp_input [3]string
    fmt.Println("Do not use space for input. Type `quit` to close program.")

    for true {
        fmt.Print("Input username: ")
        fmt.Scan(&tmp_input[0])

        if tmp_input[0] == "quit" {
            break
        }

        fmt.Print(" -> Input first name: ")
        fmt.Scan(&tmp_input[1])

        fmt.Print(" -> Input last name: ")
        fmt.Scan(&tmp_input[2])
        fmt.Println()

        // Code below will trigger error
        // --> panic: assignment to entry in nil map
        // If we do not initialize with make() first
        users[tmp_input[0]] = make(map[string]string)
        users[tmp_input[0]]["first_name"] = tmp_input[1]
        users[tmp_input[0]]["last_name"] = tmp_input[2]
    }

    fmt.Println("\nUsers list")
    fmt.Println("----------")
    for username, data := range users {
        fmt.Printf("%s: %s %s\n", username, data["first_name"], data["last_name"])
    }
}