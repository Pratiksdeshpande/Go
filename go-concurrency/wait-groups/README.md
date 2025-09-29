# WaitGroups in Go

## Introduction to WaitGroups

### What is a WaitGroup?

A WaitGroup in Go is part of the sync package.
It is used to wait for a collection of goroutines to finish executing before moving forward in your program.

Think of it as a counter:

You increase the counter when you start goroutines.

Each goroutine decreases the counter when it’s done.

The main program waits until the counter becomes zero.

### Why do we need WaitGroups?

Normally, goroutines run asynchronously. If you start multiple goroutines in main(), the main() function might exit before the goroutines finish their work.

WaitGroups help you ensure that the main program waits for all goroutines to complete before proceeding.

### The sync.WaitGroup Type

The sync package provides the WaitGroup type, which has three main methods:

1. Add(int): Increases the WaitGroup counter by the specified number.
2. Done(): Decreases the WaitGroup counter by one. This is typically called at the end of a goroutine.
3. Wait(): Blocks until the WaitGroup counter is zero.

## Basic Usage of WaitGroups

### How this works step by step:

Declare a sync.WaitGroup (like a counter).

Before starting each goroutine, call wg.Add(1).

Inside the goroutine, use defer wg.Done() → decreases the counter by 1 when the function exits.

In main(), call wg.Wait() → blocks until counter = 0.

### Key Points

Always call Add(1) before starting a goroutine.

Always use Done() inside the goroutine (best practice: defer wg.Done()).

Wait() should be called after launching all goroutines.

If Add() and Done() calls don’t match, the program will hang (Wait forever) or panic (negative counter).

## Internal Working of WaitGroups 

The Go source code defines WaitGroup (simplified) as:
```go
type WaitGroup struct {
    noCopy noCopy    // prevents copying of WaitGroup
    state1 uint64    // stores counter and waiter info
    sema   uint32    // semaphore used for blocking
}
```

Let’s break this down:

1. Counter (state1 high bits) – how many goroutines are being waited on.

2. Waiter count (state1 low bits) – how many goroutines are currently waiting (blocked on Wait()).

3. Semaphore (sema) – used to block and wake up goroutines efficiently.

### Lifecycle of a WaitGroup

#### Initialization

A WaitGroup is just a struct, no need for explicit init (value starts at zero). <br>

#### Adding (Add(n))

Increases the counter by n. <br>
Must be called before starting the goroutine.

#### Waiting (Wait())

If counter > 0 → blocks current goroutine. <br>
Uses semaphore internally to suspend goroutines instead of busy waiting.

#### Done (Done())

Decreases counter by 1.
If counter == 0 → wakes up all goroutines blocked in Wait().

### How Wait() Works

When you call wg.Wait(), the goroutine checks if the counter is zero.

If zero → it continues immediately.

If non-zero → it sleeps (blocked on semaphore).

When the last goroutine calls Done(), the counter becomes zero → all blocked goroutines waiting on semaphore are released.

## Common Use Cases

## Best Practices and Gotchas

## WaitGroups vs Other Synchronization Primitives

## Interview-style Questions and Examples