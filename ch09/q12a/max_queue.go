package ch09q12a

// This packages provides a Queue API with the standard queue operations of
// Enqueue and Dequeue as well as a Max() operation.
type Queue struct {
	queue    []int
	maxQueue []int
}

func (q *Queue) Dequeue() (val int, ok bool) {
	if len(q.queue) < 1 {
		return
	}

	val, q.queue = q.queue[0], q.queue[1:]
	if val == q.maxQueue[0] {
		q.maxQueue = q.maxQueue[1:]
	}
	ok = true
	return
}

func (q *Queue) Enqueue(val int) {
	q.queue = append(q.queue, val)

	if len(q.maxQueue) < 1 || q.maxQueue[0] < val {
		q.maxQueue = []int{val}
		return
	}

	origMaxQueueLen := len(q.maxQueue)
	for i := origMaxQueueLen - 1; i >= 0; i-- {
		if q.maxQueue[i] < val {
			q.maxQueue = q.maxQueue[:i]
		} else {
			q.maxQueue = append(q.maxQueue, val)
			break
		}
	}
}

func (q *Queue) Max() (val int, ok bool) {
	if len(q.maxQueue) < 1 {
		return
	}

	return q.maxQueue[0], true
}
