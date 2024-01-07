# If Else
This Go code demonstrates the use of conditional statements (`if` and `else`) and the use of the short declaration (`:=`). Let's go through each part of the code:

```go
package main

import "fmt"

func main() {
    // Check if 7 is even or odd
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // Check if 8 is divisible by 4
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // Check if either 8 or 7 is even
    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }

    // Using the short declaration to declare a variable 'num' and check its value
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
```

### Output
```
7 is odd
8 is divisible by 4
either 8 or 7 are even
9 has 1 digit
```

### Explanation:

1. `if 7%2 == 0 {...} else {...}`: Checks if the remainder of dividing 7 by 2 is equal to 0. If true, it prints "7 is even," otherwise, it prints "7 is odd."

2. `if 8%4 == 0 {...}`: Checks if 8 is divisible by 4 (i.e., the remainder is 0). If true, it prints "8 is divisible by 4."

3. `if 8%2 == 0 || 7%2 == 0 {...}`: Checks if either 8 or 7 is even (i.e., the remainder is 0 when divided by 2). If true, it prints "either 8 or 7 are even."

4. `if num := 9; num < 0 {...} else if num < 10 {...} else {...}`: Uses the short declaration to declare a variable `num` and initialize it with the value 9. It then checks various conditions:
   - If `num` is less than 0, it prints "9 is negative."
   - If `num` is less than 10 (but not negative), it prints "9 has 1 digit."
   - If neither condition is true, it prints "9 has multiple digits."

These conditional statements help control the flow of the program based on the given conditions. The short declaration is often used in Go for concise variable declaration and initialization within a block.
