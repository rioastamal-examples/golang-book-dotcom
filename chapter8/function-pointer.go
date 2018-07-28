package main
import "fmt"

/**
 * A pointer is indicated with * char.
 */
func changeName(name *string, newVal string) {
    // When assigning a pointer value we also still need
    // to use * char. It is called "dereference"
    *name = newVal
}

func main() {
    var myName string = "Rio Astamal"
    fmt.Println("Original Name:", myName)

    // When passing a pointer we should use &
    changeName(&myName, "Rio Astamal, S.Kom")
    fmt.Println("New Name:", myName)
}