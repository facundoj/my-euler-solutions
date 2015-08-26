// Lattice paths
// Problem 15
// Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.
//
// How many such routes are there through a 20×20 grid?

package main

import "fmt"

func main() {
    const size int = 21
    var grid [size][size]int

    for i := 0; i < size; i++ {
        if i == 0 {
            for j := 0; j < size; j++ {
                grid[0][j] = 1
                grid[j][0] = 1
            }
        } else {
            for j := i; j < size; j++ {
                grid[i][j] = grid[i - 1][j] + grid[i][j - 1]
                grid[j][i] = grid[j - 1][i] + grid[j][i - 1]
            }
        }
    }
    fmt.Println(grid[size - 1][size - 1])
}
