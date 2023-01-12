package nextiva

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

// A Logger represents an active logging object that generates lines of
// output to an io.Writer. Each logging operation makes a single call to
// the Writer's Write method. A Logger can be used simultaneously from
// multiple goroutines; it guarantees to serialize access to the Writer.
type Logger interface {

	// Log takes any value and writes it to the output.
	Log(v ...interface{})

	// SetOutput is a setter method. it sets the writer of the logger.
	SetOutput(io.Writer)
}

type ConsoleLogger struct {
	Writter io.Writer
	Mu      sync.Mutex
}

func (cl ConsoleLogger) Log(v ...interface{}) {
	cl.Mu.Lock()
	fmt.Fprint(cl.Writter, v...)
	defer cl.Mu.Unlock()
}

func (cl ConsoleLogger) SetOutput(writer io.Writer) {
	cl.Mu.Lock()
	cl.Writter = writer
	defer cl.Mu.Unlock()
}

// TODO: implement Logger
func LoggerTest(l Logger) error {
	var out bytes.Buffer
	l.SetOutput(&out)

	msg := "foo"
	l.Log(msg)

	if out.String() != msg {
		return fmt.Errorf("failed to log message")
	}
	return nil
}
