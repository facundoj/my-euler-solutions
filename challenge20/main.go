// Factorial digit sum
// Problem 20
// n! means n × (n − 1) × ... × 3 × 2 × 1
//
// For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
//
// Find the sum of the digits in the number 100!

package main

import "fmt"
import "strconv"
import "math/big"

func main()  {
    var i int64
    prod := big.NewInt(1)

    for i = 1; i <= 100; i += 1 {
        prod.Mul(prod, big.NewInt(i))
    }

    // to string
    num := fmt.Sprintf("%d", prod)

    sum := 0
    for _, charChode := range(num) {
        char := string(charChode)
        digit, _ := strconv.Atoi(char)
        sum += digit
    }

    fmt.Println(sum)
}
