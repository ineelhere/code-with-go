This Go code demonstrates the use of the `select` statement to handle multiple channel operations concurrently. Let's go through it with inline comments:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Creating two unbuffered channels, 'c1' and 'c2'
    c1 := make(chan string)
    c2 := make(chan string)

    // Launching two goroutines to send messages to 'c1' and 'c2' after a specific time delay
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    // Looping twice to handle messages from 'c1' and 'c2'
    for i := 0; i < 2; i++ {
        // The 'select' statement allows the program to wait on multiple communication operations
        // It will execute the first case that is ready, blocking the others
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```
### Output
```
received one
received two
Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.

real    0m2.245s
```
Explanation:

1. `package main`: Indicates that this Go file belongs to the main executable package.

2. `import (...)`: Imports necessary packages, including "fmt" for formatting and printing, and "time" for handling time-related operations.

3. `func main() { ... }`: The main function, where the execution of the program begins.

4. `c1 := make(chan string)`: Creates an unbuffered channel named 'c1'.

5. `c2 := make(chan string)`: Creates another unbuffered channel named 'c2'.

6. Two goroutines are launched using anonymous functions and the `go` keyword. These goroutines will send messages to 'c1' and 'c2' after specific time delays using `time.Sleep`.

7. The `for` loop runs twice to handle messages from both 'c1' and 'c2'.

8. `select { ... }`: The `select` statement allows the program to wait on multiple communication operations. It blocks until one of its cases can execute, at which point it will execute that case.

9. `case msg1 := <-c1:`: If a message is received from 'c1', it prints "received" along with the received message.

10. `case msg2 := <-c2:`: If a message is received from 'c2', it prints "received" along with the received message.

The use of `select` here allows the program to handle multiple channels concurrently, effectively waiting for the first one to send a message. The output will depend on which goroutine completes its work first.