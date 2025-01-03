# Concurrency

Remember: parallelism and concurrency are not the same thing.

## Goroutines

### WaitGroup

1. `Add(goroutines_counter)`
   1. _goroutines_counter_ = Number of goroutines that you will add
2. `Wait()`
3. `Done()`
   1. Tells the *[WaitGroup]* when a *[goroutine]* has finished.
   2. Done decrements the *[WaitGroup]* counter by one.

### Defer

## Channels

### Close

### Range

### Select

## Resources

1. [Goroutines](https://go.dev/tour/concurrency/1)