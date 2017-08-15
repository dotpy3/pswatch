# Process Watch

[![GoDoc](https://godoc.org/github.com/dotpy3/pswatch?status.svg)](https://godoc.org/github.com/dotpy3/pswatch)

A **very** simple Go package to start processes, and to be alerted when they die. Inspired by [ps-watcher](http://ps-watcher.sourceforge.net/ps-watcher.html) and [node-process-watch](https://github.com/samuelg/node-process-watch).

## Example

```go
import (
    "context"
    "fmt"
    "os"
    "time"

    "github.com/dotpy3/pswatch"
)

func main() {
    info, _ := pswatch.StartProcess(context.Background(), "sleep", []string{"5"}, &os.ProcAttr{}, time.Second)
    <-info
    fmt.Println("Sleep process is over")
}
```
