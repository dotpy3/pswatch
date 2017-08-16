# Process Watch

[![GoDoc](https://godoc.org/github.com/dotpy3/pswatch?status.svg)](https://godoc.org/github.com/dotpy3/pswatch)

A **very** simple Go package to watch processes, and to be alerted when they end. Inspired by [ps-watcher](http://ps-watcher.sourceforge.net/ps-watcher.html) and [node-process-watch](https://github.com/samuelg/node-process-watch).

## Example

```go
import (
    "context"
    "fmt"

    "github.com/dotpy3/pswatch"
)

func watchPID(pid int) {
    // Start a process, without being subject to a context, and poll it every second for status
    info, _ := pswatch.WatchProcess(context.Background(), pid, pswatch.DefaultPollMargin)
    <-info
    fmt.Println("Sleep process is over")
}
```
