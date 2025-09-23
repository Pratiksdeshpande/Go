# 🚀 Goroutines in Go – Complete Guide

## 🔹 What is a Goroutine?

A **goroutine** is a lightweight, independently executing function managed by the Go runtime.

**Created using the `go` keyword:**
```go
go myFunction()
```

### Key Features:
- 🪶 **Lightweight**: Unlike OS threads, goroutines are cheap to create and manage (initial stack ~2 KB)
- 📈 **Scalable**: You can spawn thousands or even millions of goroutines
- ⚡ **Fast**: Managed entirely in user space for efficiency

> ⚠️ **Important**: If the main goroutine exits, all child goroutines are killed immediately.

---

## 🔹 How are Goroutines Scheduled?

The Go runtime uses an **M:N scheduler**:

- **M** = OS threads (machines)
- **N** = Goroutines multiplexed on top of M threads
- **Managed entirely in user space** → very efficient

### Key Benefits:
- ➡️ Goroutines are scheduled across multiple OS threads without kernel intervention
- ➡️ Uses work-stealing to balance load between processors

---

## 🔹 How Goroutines Differ from Threads?

| Feature | 🟢 Goroutine | 🔴 OS Thread |
|---------|--------------|--------------|
| **Managed by** | Go runtime (user space) | OS kernel |
| **Stack size** | ~2 KB (dynamic growth/shrink) | ~1 MB (fixed, large upfront) |
| **Context switch** | User space (fast) | Kernel space (slow) |
| **Creation cost** | Very cheap | Expensive |
| **Scalability** | Millions possible | Thousands at best |
| **Scheduling** | Go runtime scheduler | OS kernel scheduler |

---

## 🔹 Scheduler – G, M, P Model

The Go scheduler works on **three entities**:

### 🔄 Components:
- **G (Goroutine)** → execution context (function, stack, instruction pointer)
- **M (Machine)** → an OS thread that executes goroutines
- **P (Processor)** → provides context for scheduling (run queues)

### 📊 Flow:
1. Goroutines (G) are pushed into P's local queue
2. M (thread) bound to a P executes goroutines from that queue
3. If P's queue is empty, it can steal goroutines from another P's queue

> 👉 **The number of P = GOMAXPROCS** (defaults to number of CPU cores)

```go
#Go

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}
```

---

## 🔹 Goroutine Stack Memory & Lifecycle

### 📦 Stack Memory:
- **Initial size**: ~2 KB
- **Dynamic growth**: Grows when deeper recursion or larger data is needed
- **Smart shrinking**: Shrinks back when not required
- **High capacity**: Enables millions of goroutines in a program

> ⚠️ **Warning**: Creating too many goroutines may still cause memory pressure → be careful with unbounded spawning.

---

## 🔹 Goroutines & Garbage Collector (GC)

### 🗑️ Memory Management:
- Each goroutine's stack is allocated on the heap
- When goroutine exits → stack is garbage-collected
- If a goroutine is blocked forever (goroutine leak), its memory is never freed

### 💡 Best Practices:
- ✅ Always ensure goroutines have a termination condition (use `context.Context`)
- 📊 Monitor with `runtime.NumGoroutine()` in debugging or `pprof` in production

---

## 🔹 Complete Goroutine Lifecycle

### 🔄 Lifecycle States:

1. **🆕 Created** → with `go func() {...}`
2. **⏳ Runnable** → placed in P's run queue
3. **▶️ Running** → picked by an M (thread) bound to a P
4. **⏸️ Waiting** → blocked on:
   - 📡 Channel operation
   - 🌐 Network / I/O
   - 🔒 Synchronization primitives (Mutex, WaitGroup, etc.)
5. **✅ Terminated** → function finishes → goroutine is cleaned up by GC