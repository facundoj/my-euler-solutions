package main

import "fmt"
import "strconv"
import "math/big"

func main() {
    num := big.NewInt(2)
    num.Exp(num, big.NewInt(1000), nil)

    // to string
    numStr := fmt.Sprintf("%d", num)

    sum := 0
    for _, charChode := range(numStr) {
        char := string(charChode)
        digit, _ := strconv.Atoi(char)
        sum += digit
    }

    fmt.Println(sum)
}
