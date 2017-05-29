# concurrency

A showcase of concurrency design patters in go that i find useful.

## Dispatcher

Several workers process jobs in parallel. A dispatcher distributes tasks among the workers and prevents the incoming queue from filling up.

### Advantages
* Many workers
* Does not block producers

### Drawbacks
* Starvation?

## Burst

A way of rate limiting incoming requests. Currently only works for a single receiver.

### Drawbacks
* Single worker
