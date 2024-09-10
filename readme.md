# 7 ways to profile Golang Application

1. `time go fmt std` on the cmd with the file you are evaluating

2. `env GODEBUG=gctrace=1 godoc -http-:8080'

## types of profiling

- CPU profiling - involves capturing detailed information about how your program utilizes the CPU over time. It provides insights into which functions consume the most CPU time, helping developers identify performance bottlenecks and optimize their code.
- Memory profiling - records the stack trace when a heap allocation is made
- Block profiling - records the amount of time a goroutine spent waiting for a shared resource


`import (
    _ "net/http/pprof"
    "net/http"
)

func main() {
    http.ListenAndServe(":8080", nil)
}` 

## Analyzing CPU Profile Data

`go tool pprof cpu.prof`

Common Commands:
- top: Displays the top functions by CPU usage.
- list <function>: Shows the source code of a specific function with annotated CPU usage.
- web: Generates and opens a graphical representation of the profile in a web browser.
- svg: Creates an SVG file with a visual representation of the CPU usage.