package sfp

import (
	"log"
	"sync"
)

type RingBuffer struct {
	Buff     [][]byte
	capacity int
	head     int
	tail     int
	size     int
	mu       sync.Mutex
}

func NewRingBuffer(capacity int) *RingBuffer {
	return &RingBuffer{
		Buff:     make([][]byte, capacity),
		capacity: capacity,
		head:     0, //Pointer on place to Pop()
		tail:     0, //Increases on Push()
		size:     0, //Current buffer size
		mu:       sync.Mutex{},
	}
}

func (rb *RingBuffer) Size() int {
	return rb.size
}

func (rb *RingBuffer) Tail() int {
	return rb.tail
}

func (rb *RingBuffer) Push(buff []byte) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if rb.size == rb.capacity {

		rb.head = (rb.head + 1) % rb.capacity

	} else {

		rb.size++
	}

	rb.Buff[rb.tail] = buff

	rb.tail = (rb.tail + 1) % rb.capacity
	log.Printf("tail: %d\n", rb.tail)
	log.Printf("head: %d\n", rb.head)
}

func (rb *RingBuffer) Pop() []byte {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if rb.size == 0 {
		return nil
	}

	data := rb.Buff[rb.head]
	rb.head = (rb.head + 1) % rb.capacity
	rb.size--

	return data
}
