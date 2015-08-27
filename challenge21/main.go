// Amicable numbers
// Problem 21
// Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
// If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.
//
// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
//
// Evaluate the sum of all the amicable numbers under 10000.

package main

import "fmt"

func main() {
    amicables := make(map[int]bool)

    for i := 1; i <= 10000; i += 1 {
        posiblyAmic := sumAll(getDivisors(i))

        if i != posiblyAmic && i == sumAll(getDivisors(posiblyAmic)) {
            fmt.Printf("%d and %d\n", i, posiblyAmic)
            amicables[i] = true
            amicables[posiblyAmic] = true
        }
    }

    amicablesSum := 0
    for num, _ := range amicables {
        amicablesSum += num
    }

    fmt.Println(amicablesSum)
}

func sumAll(list []int) int {
    sum := 0
    for _, num := range list {
        sum += num
    }

    return sum
}

func getDivisors(num int) []int {
    divisors := make([]int, num)

    for i := 1; i < num; i += 1 {
        if num % i == 0 {
            divisors = append(divisors, i)
        }
    }

    return divisors
}
