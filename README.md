This is one of the first packages I wrote in Go.

I wanted to have complete control over how a Logger prefixes its output, so I wrote this.

Here's an example:

```go
package main

import (
	"os"
	"time"

	"github.com/zaydek/logger"
)

func main() {
	l := logger.New(os.Stdout, func() string { return time.Now().Format("2006-01-02 15:04:05") })
	l.Println("hello, world!") // e.g. "2006-01-02 15:04:05 hello, world!"
}
```

Here's a more granular example:

```go
func main() {
	l := logger.New(os.Stderr, func() string { return time.Now().Format("15:04:05.000000") }) // usecs
	l.Println("hello, darkness...") // e.g. "15:04:05.000000 hello, darkness..."
}
```
