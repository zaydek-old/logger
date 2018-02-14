package logger

import (
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	m sync.Mutex
	w io.Writer
	f func() string
}

func New(w io.Writer, f func() string) *Logger {
	return &Logger{f: f, w: w}
}

func (l *Logger) Printf(format string, a ...interface{}) {
	l.m.Lock()
	fmt.Fprintf(l.w, format, prepend(l.f(), a)...)
	l.m.Unlock()
}

func (l *Logger) Println(a ...interface{}) {
	l.m.Lock()
	fmt.Fprintln(l.w, prepend(l.f(), a)...)
	l.m.Unlock()
}

func prepend(v interface{}, a []interface{}) []interface{} {
	return append([]interface{}{v}, a...)
}
