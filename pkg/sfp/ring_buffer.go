package sfp

import (
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

	if rb.size == rb.capacity {

		rb.head = (rb.head + 1) % rb.capacity

	} else {

		rb.size++
	}

	rb.Buff[rb.tail] = buff

	rb.tail = (rb.tail + 1) % rb.capacity
	rb.mu.Unlock()
}

func (rb *RingBuffer) Pop() []byte {
	rb.mu.Lock()

	if rb.size == 0 {
		rb.mu.Unlock()
		return nil
	}

	data := rb.Buff[rb.head]
	rb.head = (rb.head + 1) % rb.capacity
	rb.size--

	rb.mu.Unlock()
	return data
}
