Fifo Channel Queue (String)
===========================

## Purpose:
* Provide a blocking FIFO queue that ensures strings sent to the
  queue are persisted.  If the queue is at capacity, the `.Push()` 
  method should block until the queue length has decreased.

## Use Case:
Where we may provision a buffered string channel to hold up to 
n elements, we may have an asynchronous generator pushing new data
into the queue faster than the consumer is popping elements from the
other end.  By default, this would cause data to be discarded from
the channel, losing data integrity.

To avoid the problem, Fifo::Push() will block indefinitely if the 
channel is at capacity, allowing the receiver to consume elements
from the queue.

## Methods
`.Setup(sz)` - Initialize the channel of a given size (sz) for use.

`.Pop()`     - Pop a message off the queue.

`.Push(s)`   - Push a string (s) to the message queue, blocking the
               push operation if the queue is at capacity until
               messages are popped from the queue.