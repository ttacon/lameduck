package lameduck

import (
	"os"
	"os/signal"
)

type LameDuckHandler interface {
	// Will enter lame duck mode when the given signal is received.
	WithSignalHandler(signal os.Signal) LameDuckHandler
	// A convenience method for entering lameduck mode on SIGINT (2).
	WithSigINTHandler() LameDuckHandler
	// Launches all the registered listeners for lame duck mode.
	Go()
}

type EnterLameDuckMode func()

func NewLameDuckHandler(fn EnterLameDuckMode) LameDuckHandler {
	return &lameDuckHandler{
		fn:      fn,
		entered: make(chan struct{}),
	}
}

type lameDuckHandler struct {
	fn        EnterLameDuckMode
	listeners []func()
	entered   chan struct{}
}

func (l *lameDuckHandler) WithSignalHandler(signal os.Signal) LameDuckHandler {
	l.listeners = append(l.listeners, l.signalHandler(signal))
	return l
}

func (l *lameDuckHandler) signalHandler(s os.Signal) func() {
	return func() {
		go func() {
			c := make(chan os.Signal, 1)
			signal.Notify(c, s)
			select {
			case <-c:
				l.enterLameDuckMode()
			case <-l.entered:
				signal.Stop(c)
			}
		}()
	}
}

func (l *lameDuckHandler) WithSigINTHandler() LameDuckHandler {
	return l.WithSignalHandler(os.Interrupt)
}

func (l *lameDuckHandler) Go() {
	// launch all the handlers
	for _, fn := range l.listeners {
		// TODO(ttacon): use timeouts to know if this function isn't returning
		fn()
	}
}

func (l *lameDuckHandler) enterLameDuckMode() {
	// shut down all other handlers
	close(l.entered)

	l.fn()
}

// TODO(ttacon): add HTTP handler
// TODO(ttacon): add TCP handler
// TODO(ttacon): other lame listeners?
