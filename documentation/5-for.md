# For Loop

This Go code demonstrates various usage of the `for` loop, including basic loop structure, loop with a condition, an infinite loop with a `break` statement, and using `continue` to skip certain iterations. Let's go through each part of the code:

```go
package main

import "fmt"

func main() {
    // Basic for loop with a condition
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // For loop with an initialization statement, a condition, and a post statement
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // Infinite loop with a break statement
    for {
        fmt.Println("loop")
        break
    }

    // For loop with continue statement to skip even numbers
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
```
### Output
```
1
2
3
7
8
9
loop
1
3
5
```
### Explanation:

1. `for i <= 3 {...}`: A basic `for` loop with a condition. It initializes `i` to 1 and continues looping as long as `i` is less than or equal to 3. It prints the value of `i` in each iteration and increments `i` by 1.

2. `for j := 7; j <= 9; j++ {...}`: Another `for` loop with an initialization statement (`j := 7`), a condition (`j <= 9`), and a post statement (`j++`). It initializes `j` to 7, continues looping as long as `j` is less than or equal to 9, prints the value of `j` in each iteration, and increments `j` by 1.

3. `for {...}`: An infinite loop. It continually prints "loop" and breaks out of the loop using the `break` statement after the first iteration.

4. `for n := 0; n <= 5; n++ {...}`: A `for` loop with an initialization statement, a condition, and a post statement. It initializes `n` to 0, continues looping as long as `n` is less than or equal to 5, increments `n` by 1 in each iteration, and prints the value of `n` only if it's an odd number (using `if n%2 == 0 { continue }` to skip even numbers).
