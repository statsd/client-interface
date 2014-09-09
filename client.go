package statsd

import "time"

type Client interface {
	Gauge(name string, value int) error

	Incr(name string) error
	IncrBy(name string, count int) error

	Decr(name string) error
	DecrBy(name string, count int) error

	Duration(name string, duration time.Duration) error
	Histogram(name string, value int) error

	Annotate(name string, fmt string, args ...interface{}) error

	Flush() error
}
