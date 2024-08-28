package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

var globalCloser = New()

// Add adds new closers
func Add(fs ...func() error) {
	globalCloser.Add(fs...)
}

// Wite wites all closers
func Wite() {
	globalCloser.Wite()
}

// CloseAll closes all closers functions
func CloseAll() {
	globalCloser.CloseAll()
}

// Closer closes all closers
type Closer struct {
	done chan struct{}
	once sync.Once
	mu   sync.RWMutex
	fs   []func() error
}

// New creates a new instance of Closer
func New(signals ...os.Signal) *Closer {
	c := &Closer{done: make(chan struct{})}

	if len(signals) > 0 {
		go func() {
			sign := make(chan os.Signal, 1)
			signal.Notify(sign, signals...)
			<-sign
			signal.Stop(sign)

			c.CloseAll()
		}()

	}

	return c
}

// Add adds new closers
func (c *Closer) Add(fs ...func() error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.fs = append(c.fs, fs...)
}

// Wite wites all closers
func (c *Closer) Wite() {
	<-c.done
}

// CloseAll closes all closers functions
func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()
		fs := c.fs
		c.fs = nil
		c.mu.Unlock()

		errs := make(chan error, len(fs))

		for _, f := range fs {
			go func(f func() error) {
				errs <- f()
			}(f)
		}

		for i := 0; i < cap(errs); i++ {
			if err := <-errs; err != nil {
				log.Println("error returned from Closer")
			}
		}
	})
}
