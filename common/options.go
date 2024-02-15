package common

type Options struct {
	Debug     bool
	File      string
	Goroutine int64
	Output    string
	Timeout   int64
}

var Opts Options
