This Go code demonstrates the use of a combination of channels and goroutines to synchronize the sending and receiving of messages. Let's go through it with inline comments:

```go
package main

import "fmt"

func main() {
    // Creating a buffered channel 'jobs' with a capacity of 5
    jobs := make(chan int, 5)

    // Creating an unbuffered channel 'done' for signaling when all jobs are received
    done := make(chan bool)

    // Launching a goroutine to receive jobs from the 'jobs' channel
    go func() {
        for {
            // Attempting to receive a job from 'jobs'
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                // If 'more' is false, it means the 'jobs' channel is closed
                fmt.Println("received all jobs")
                // Signaling that all jobs are received by sending true to 'done'
                done <- true
                return
            }
        }
    }()

    // Sending three jobs to the 'jobs' channel
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }

    // Closing the 'jobs' channel to indicate that no more jobs will be sent
    close(jobs)
    fmt.Println("sent all jobs")

    // Waiting for the 'done' channel to receive a signal indicating all jobs are received
    <-done

    // Attempting to receive from the closed 'jobs' channel
    _, ok := <-jobs
    fmt.Println("received more jobs:", ok)
}
```
### Output
```
sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs
received more jobs: false
```
Explanation:

1. `package main`: Indicates that this Go file belongs to the main executable package.

2. `import "fmt"`: Imports the "fmt" package for formatting and printing.

3. `func main() { ... }`: The main function, where the execution of the program begins.

4. `jobs := make(chan int, 5)`: Creates a buffered channel named 'jobs' with a capacity of 5.

5. `done := make(chan bool)`: Creates an unbuffered channel named 'done' for signaling when all jobs are received.

6. A goroutine is launched to receive jobs from the 'jobs' channel. It loops indefinitely, attempting to receive jobs until the 'jobs' channel is closed.

7. Inside the goroutine, `more` is a boolean value that indicates whether there are more values to receive from the 'jobs' channel.

8. The main goroutine sends three jobs to the 'jobs' channel and prints a message for each sent job.

9. `close(jobs)`: Closes the 'jobs' channel to indicate that no more jobs will be sent.

10. `<-done`: Waits for the 'done' channel to receive a signal (true) indicating that all jobs are received.

11. `_, ok := <-jobs`: Attempts to receive from the closed 'jobs' channel. Since the channel is closed, `ok` will be false, and it prints "received more jobs: false."

In summary, this code demonstrates how to use channels and goroutines to coordinate the sending and receiving of messages, and how to signal the completion of tasks using channels. The `close` operation on the 'jobs' channel is used to signal that no more jobs will be sent.