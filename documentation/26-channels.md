This Go code demonstrates the use of channels for communication between goroutines. Let's break it down with inline comments:

```go
package main

import "fmt"

func main() {
	// Creating a channel named 'messages' for communication between goroutines
	messages := make(chan string)

	// Launching a goroutine with an anonymous function to send a message into the channel
	go func() {
		messages <- "ping" // Sending "ping" into the 'messages' channel
	}()

	// Receiving the message from the channel and storing it in the variable 'msg'
	msg := <-messages

	// Printing the received message
	fmt.Println(msg)
}
```
### Output
```
ping
```

Explanation:

1. `package main`: Indicates that this Go file belongs to the main executable package.

2. `import "fmt"`: Imports the "fmt" package for formatting and printing.

3. `func main() { ... }`: The main function, where the execution of the program begins.

4. `messages := make(chan string)`: Creates a channel named 'messages' for communication between goroutines. The type of data that can be sent through this channel is a string.

5. `go func() { messages <- "ping" }()`: Launches a new goroutine with an anonymous function. This function sends the string "ping" into the 'messages' channel. This is an example of sending data to a channel.

6. `msg := <-messages`: Receives a message from the 'messages' channel and assigns it to the variable 'msg'. This is an example of receiving data from a channel.

7. `fmt.Println(msg)`: Prints the received message, which is "ping" in this case.

In summary, this code demonstrates the basic usage of channels for communication between goroutines. The main goroutine creates a channel, launches another goroutine to send a message into the channel, and then receives and prints the message in the main goroutine. This showcases how channels facilitate communication and synchronization between concurrently running goroutines.