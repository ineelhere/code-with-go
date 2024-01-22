This Go code demonstrates a simple worker pool pattern where multiple worker goroutines process jobs concurrently. Let's go through it with inline comments:

```go
package main

import (
    "fmt"
    "time"
)

// worker function represents a worker that processes jobs
func worker(id int, jobs <-chan int, results chan<- int) {
    // Loop over jobs received from the 'jobs' channel
    for j := range jobs {
        fmt.Println("worker", id, "started job", j)
        // Simulate work by sleeping for one second
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        // Send the result (j * 2) to the 'results' channel
        results <- j * 2
    }
}

func main() {
    const numJobs = 5

    // Creating two channels: 'jobs' for sending jobs to workers, and 'results' for receiving results
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Launching three worker goroutines
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Sending five jobs to the 'jobs' channel
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    
    // Closing the 'jobs' channel to indicate that no more jobs will be sent
    close(jobs)

    // Receiving results from the 'results' channel
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
```
### Output
```
worker 3 started  job 1
worker 2 started  job 3
worker 1 started  job 2
worker 1 finished job 2
worker 1 started  job 4
worker 2 finished job 3
worker 3 finished job 1
worker 2 started  job 5
worker 2 finished job 5
worker 1 finished job 4
```
Explanation:

1. `package main`: Indicates that this Go file belongs to the main executable package.

2. `import (...)`: Imports necessary packages, including "fmt" for formatting and printing, and "time" for handling time-related operations.

3. `func worker(id int, jobs <-chan int, results chan<- int) { ... }`: Defines a worker function that processes jobs received from the 'jobs' channel and sends results to the 'results' channel.

4. `for j := range jobs { ... }`: The worker goroutine loops over jobs received from the 'jobs' channel. The loop continues until the 'jobs' channel is closed.

5. `jobs <- j`: Sends the current job to the worker goroutine.

6. `time.Sleep(time.Second)`: Simulates work by sleeping for one second.

7. `results <- j * 2`: Sends the result (j * 2) to the 'results' channel.

8. `func main() { ... }`: The main function, where the execution of the program begins.

9. `jobs := make(chan int, numJobs)`: Creates a buffered channel named 'jobs' for sending jobs to worker goroutines. The buffer size is set to the number of jobs to allow non-blocking sends.

10. `results := make(chan int, numJobs)`: Creates a buffered channel named 'results' for receiving results from worker goroutines.

11. `for w := 1; w <= 3; w++ { go worker(w, jobs, results) }`: Launches three worker goroutines, each with its own ID, 'jobs' channel for receiving jobs, and 'results' channel for sending results.

12. `for j := 1; j <= numJobs; j++ { jobs <- j }`: Sends five jobs to the 'jobs' channel.

13. `close(jobs)`: Closes the 'jobs' channel to indicate that no more jobs will be sent.

14. `for a := 1; a <= numJobs; a++ { <-results }`: Receives and discards results from the 'results' channel. This loop ensures that all results are processed.

This code demonstrates the basic concept of a worker pool pattern in Go, where multiple worker goroutines concurrently process jobs. The worker goroutines are launched as separate goroutines, and jobs are sent to them via a channel. Results are collected through another channel. The program waits for all results to be processed before exiting.