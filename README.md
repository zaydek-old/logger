# logger

This is one of the first packages I wrote in Go.

I was frustrated with the lack of granular controls `pkg/log` predefines in its constants, e.g. `log.Ltime`, etc., so I wrote this, instaed.

Here's example usage:

```
package main

import (
	"os"
	"time"

	"owl.delivery/etc/logger"
)

func main() {
	l := logger.New(os.Stdout, func() string { return time.Now().Format("2006-01-02 15:04:05") })
	l.Println("hello, world!") // e.g. "2006-01-02 15:04:05 hello, world!"
}
```

And a more granular example:

```
func main() {
	l := logger.New(os.Stderr, func() string { return time.Now().Format("15:04:05.000000") }) // usecs
	l.Println("hello, darkness...") // e.g. "15:04:05.000000 hello, darkness..."
}
```
