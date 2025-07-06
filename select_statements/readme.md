# Select Statement Examples

## 1. Multiplexing Receives

Wait on multiple incoming channels and handle whichever sends first.

```go
package main

import (
    "fmt"
    "time"
)

func ping(ch chan<- string) {
    time.Sleep(500 * time.Millisecond)
    ch <- "ping"
}

func pong(ch chan<- string) {
    time.Sleep(300 * time.Millisecond)
    ch <- "pong"
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go ping(ch1)
    go pong(ch2)

    // select picks whichever channel is ready
    select {
    case msg := <-ch1:
        fmt.Println("Received from ch1:", msg)
    case msg := <-ch2:
        fmt.Println("Received from ch2:", msg)
    }
}
```

Here, `pong` usually wins because it sleeps less. The program prints:

```
Received from ch2: pong
```

---

## 2. Implementing Timeouts

Combine a channel receive with a `time.After` to enforce a timeout.

```go
package main

import (
    "fmt"
    "time"
)

func longTask(done chan<- bool) {
    // Simulate a long operation
    time.Sleep(2 * time.Second)
    done <- true
}

func main() {
    done := make(chan bool)
    go longTask(done)

    select {
    case <-done:
        fmt.Println("Task completed!")
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout: gave up waiting.")
    }
}
```

If `longTask` takes more than 1 s, the timeout branch fires.

---

## 3. Non‑Blocking Operations with `default`

Use a `default` case to attempt a send or receive without blocking.

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 1)
    ch <- 42

    select {
    case val := <-ch:
        fmt.Println("Got:", val)
    default:
        fmt.Println("No value ready, moving on.")
    }

    select {
    case ch <- 7:
        fmt.Println("Sent 7 into channel")
    default:
        fmt.Println("Channel was full, could not send")
    }
}
```

With the buffered channel already holding one element, the second select’s send will go to `default`.

---

## 4. Detecting Closed Channels

When a channel is closed, receives return the zero value immediately.

```go
package main

import "fmt"

func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 3; i++ {
            ch <- i
        }
        close(ch)
    }()

    for {
        select {
        case v, ok := <-ch:
            if !ok {
                fmt.Println("Channel closed; exiting")
                return
            }
            fmt.Println("Received:", v)
        }
    }
}
```

Here, after the three sends and `close(ch)`, the `ok` flag becomes `false` and the loop exits.

---

## 5. Fan‑In Pattern

Merge multiple producer channels into one consumer channel using `select`.

```go
package main

import (
    "fmt"
    "time"
)

func producer(name string, ch chan<- string) {
    for i := 1; i <= 3; i++ {
        ch <- fmt.Sprintf("%s: %d", name, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func fanIn(chs ...<-chan string) <-chan string {
    out := make(chan string)
    for _, ch := range chs {
        go func(c <-chan string) {
            for msg := range c {
                out <- msg
            }
        }(ch)
    }
    return out
}

func main() {
    a := make(chan string)
    b := make(chan string)

    go producer("A", a)
    go producer("B", b)

    merged := fanIn(a, b)
    for i := 0; i < 6; i++ {
        fmt.Println(<-merged)
    }
}
```

The `fanIn` helper launches a goroutine per input channel, forwarding all messages to `out`.

---

## 6. Periodic Work with Tickers

Use `time.Ticker` in a `select` loop to run tasks at intervals, and stop cleanly.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(500 * time.Millisecond)
    quit := make(chan struct{})

    go func() {
        time.Sleep(2 * time.Second)
        close(quit)
    }()

    for {
        select {
        case t := <-ticker.C:
            fmt.Println("Tick at", t.Format("15:04:05.000"))
        case <-quit:
            fmt.Println("Stopping ticker")
            ticker.Stop()
            return
        }
    }
}
```

This prints ticks every 500 ms until `quit` is closed after 2 s.

---

### Summary of Key Patterns

-   **Basic multiplexing**: wait on several channels.
-   **Timeouts**: combine with `time.After`.
-   **Non-blocking**: use `default` to avoid blocking.
-   **Closure detection**: check the `ok` flag on receive.
-   **Fan‑in/fan‑out**: merge or distribute work across channels.
-   **Tickers and periodic tasks**: handle regular intervals cleanly.
