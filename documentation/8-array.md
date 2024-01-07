# Array

This Go (Golang) code is a simple program that demonstrates the usage of arrays in Go. 

Let's break down the code step by step:
```go
package main

import "fmt"

func main() {

    // Declare an array 'a' of integers with a length of 5
    var a [5]int
    fmt.Println("emp:", a) // Print the content of array 'a' (initialized with zeros)
    
    // Set the fifth element of array 'a' to 100
    a[4] = 100
    fmt.Println("set:", a) // Print the modified array 'a'
    fmt.Println("get:", a[4]) // Print the value of the fifth element in array 'a'

    fmt.Println("len:", len(a)) // Print the length of array 'a'

    // Declare and initialize an array 'b' with specific values
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b) // Print the content of array 'b'

    // Declare a 2D array 'twoD' with 2 rows and 3 columns
    var twoD [2][3]int
    
    // Populate the 2D array 'twoD' with values based on the sum of row and column indices
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    
    fmt.Println("2d: ", twoD) // Print the 2D array 'twoD'
}
```

### Output

```
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]
```

### More Explanation

```go
package main

import "fmt"
```

This code defines a Go program. The `import "fmt"` statement imports the "fmt" package, which provides functions for formatting and printing output.

```go
func main() {
```

Here begins the `main` function, which is the entry point of any Go program.

```go
    var a [5]int
    fmt.Println("emp:", a)
```

This code declares an integer array `a` of size 5. Arrays in Go have a fixed size, and in this case, it's `[5]int`. The elements of the array are initialized with their zero values (0 for integers in this case). It then prints the array using `fmt.Println`.

```go
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
```

This section demonstrates how to set and get values from the array. It sets the value at index 4 of array `a` to 100, then prints the modified array and the value at index 4.

```go
    fmt.Println("len:", len(a))
```

This line prints the length of the array `a` using the `len` function. In this case, it will print 5 because the array has 5 elements.

```go
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
```

Here, a new array `b` is declared and initialized with values in a single line. This is a shorthand syntax for array initialization.

```go
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
```

This part demonstrates the creation and initialization of a 2D array (`twoD`) with dimensions 2x3. The nested loop sets each element in the 2D array to the sum of its row and column indices. Finally, it prints the 2D array.

In summary, this Go program covers basic array operations, including declaration, initialization, setting values, getting values, finding the length of an array, and working with 2D arrays. It's a good starting point for understanding the fundamentals of arrays in Go.