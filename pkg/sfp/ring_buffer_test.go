package sfp

import (
	context2 "context"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBufferSpeed(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(2)

	rb := NewRingBuffer(1500)
	context, cancel := context2.WithTimeout(context2.Background(), 60*time.Second)
	defer cancel()

	time.Sleep(1 * time.Second)
	var counter int64

	var BenchPKG = append([]byte("HEADER_ID:12345|TIMESTAMP:99999|PAYLOAD:"), make([]byte, 98)...)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-context.Done():
				return
			default:
				rb.Push(BenchPKG)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-context.Done():
				log.Printf("results: %d package handle\n", atomic.LoadInt64(&counter))
				return
			default:
				example_data := rb.Pop()
				if example_data != nil {
					atomic.AddInt64(&counter, 1)
				} else {
					runtime.Gosched()
				}
			}
		}
	}()

	wg.Wait()
}
