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

/**
 * A new() is built-in Go function. It similar with & which
 * mean to allocate a pointer value.
 *
 * In most cases it's better to use &.
 */
func main() {
    myName := new(string)

    // Here we are not using & to pass variable because
    // it was created using "new" keyword
    changeName(myName, "Rio Astamal")

    // But we need to use * to point to pointer location
    // when printing the variable
    fmt.Println("Original Name:", *myName)

    changeName(myName, "Rio Astamal, S.Kom")
    fmt.Println("New Name:", *myName)
}