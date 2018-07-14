package main
import "fmt"

func main() {
    const MY_PREFIX string = "_RIO_"

    // Unsigned 32 bits
    const MY_NUMBER int = 19870604

    // Uncomment code below to see errors
    // MY_PREFIX = "RIO"
    // MY_NUMBER = 1000

    fmt.Println(MY_PREFIX)
    fmt.Println(MY_NUMBER)
}