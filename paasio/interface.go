package paasio

import (
	"io"
	"sync"
)

// ReadCounter is an interface describing objects that can be read from,
// and that can count the number of times they have been read from.
//
// If multiple goroutines concurrently call Read, implementations are not
// required to provide any guarantees about interleaving of the Read calls.
// However, implementations MUST guarantee that calls to ReadCount always return
// correct results even in the presence of concurrent Read calls.
type ReadCounter interface {
	io.Reader
	// ReadCount returns the total number of bytes successfully read along
	// with the total number of calls to Read().
	ReadCount() (n int64, nops int)
}

// WriteCounter is an interface describing objects that can be written to,
// and that can count the number of times they have been written to.
//
// If multiple goroutines concurrently call Write, implementations are not
// required to provide any guarantees about interleaving of the Write calls.
// However, implementations MUST guarantee that calls to WriteCount always return
// correct results even in the presence of concurrent Write calls.
type WriteCounter interface {
	io.Writer
	// WriteCount returns the total number of bytes successfully written along
	// with the total number of calls to Write().
	WriteCount() (n int64, nops int)
}

// ReadWriteCounter is the union of ReadCounter and WriteCounter.
//
// All guarantees that apply to either of ReadCounter or WriteCounter
// also apply to ReadWriteCounter.
type ReadWriteCounter interface {
	ReadCounter
	WriteCounter
}

type readCounter struct {
	r      io.Reader
	nbytes int64
	nops   int
	sync.RWMutex
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.r.Read(p)
	rc.Lock()
	rc.nops += 1
	rc.nbytes += int64(n)
	rc.Unlock()

	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.RLock()
	defer rc.RUnlock()
	return rc.nbytes, rc.nops
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{
		r: r,
	}
}

type writeCounter struct {
	w      io.Writer
	nbytes int64
	nops   int
	sync.RWMutex
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.w.Write(p)
	wc.Lock()
	wc.nops += 1
	wc.nbytes += int64(n)
	wc.Unlock()

	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.RLock()
	wc.RUnlock()
	return wc.nbytes, wc.nops
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{
		w: w,
	}
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		NewReadCounter(rw),
		NewWriteCounter(rw),
	}
}
