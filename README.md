This logger usses a string-returning func passed in at construction to _generate_ output per call, as supposed to static text. I wanted to have complete control over the details of how `pkg/log` formats time, more than it's predefined constants, so this idea came to light.

Here's an example of usage:

```package main

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

And a more granular example of usage:

```
func main() {
	l := logger.New(os.Stdout, func() string { return time.Now().Format("2006-01-02 15:04:05.000000") }) // microseconds
	l.Println("hello, world!") // e.g. "2006-01-02 15:04:05.000000 hello, world!"
}
```
