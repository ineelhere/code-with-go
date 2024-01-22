This Go code demonstrates the communication between two goroutines using channels. Let's go through it with inline comments:

```go
package main

import "fmt"

// Function ping sends a message to the provided channel
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// Function pong receives a message from one channel and sends it to another channel
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    // Creating two buffered channels, 'pings' and 'pongs', each with a capacity of 1
    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    // Sending a message to the 'pings' channel
    ping(pings, "passed message")

    // Executing the 'pong' function with the 'pings' and 'pongs' channels
    pong(pings, pongs)

    // Receiving and printing the final message from the 'pongs' channel
    fmt.Println(<-pongs)
}
```
### Output
```
passed message
```
Explanation:

1. `package main`: Indicates that this Go file belongs to the main executable package.

2. `import "fmt"`: Imports the "fmt" package for formatting and printing.

3. `func ping(pings chan<- string, msg string) { ... }`: Defines a function `ping` that sends a message (`msg`) to the provided channel (`pings`). The channel is specified as a send-only channel (`chan<- string`), meaning it can only be used for sending.

4. `func pong(pings <-chan string, pongs chan<- string) { ... }`: Defines a function `pong` that receives a message from one channel (`pings`) and sends it to another channel (`pongs`). The 'pings' channel is specified as a receive-only channel (`<-chan string`), and the 'pongs' channel is specified as a send-only channel (`chan<- string`).

5. `func main() { ... }`: The main function, where the execution of the program begins.

6. `pings := make(chan string, 1)`: Creates a buffered channel named 'pings' with a capacity of 1.

7. `pongs := make(chan string, 1)`: Creates a buffered channel named 'pongs' with a capacity of 1.

8. `ping(pings, "passed message")`: Calls the `ping` function to send the message "passed message" to the 'pings' channel.

9. `pong(pings, pongs)`: Calls the `pong` function with the 'pings' and 'pongs' channels.

10. `fmt.Println(<-pongs)`: Receives and prints the final message from the 'pongs' channel. This demonstrates the successful communication between the two goroutines.

In summary, this code illustrates how two goroutines communicate by passing a message between them using channels. The channels are used to coordinate the flow of data between concurrent parts of the program.