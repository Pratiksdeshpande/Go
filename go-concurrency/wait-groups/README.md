# 🤔 What is a WaitGroup?

---
## 📑 Table of Contents

- [Introduction to WaitGroups](#-introduction-to-waitgroups)
- [Basic Usage of WaitGroups](#-basic-usage-of-waitgroups)
- [Internal Working of WaitGroups](#-internal-working-of-waitgroups)
- [Key Points](#-key-points)
- [Best Practices](#-best-practices)
- [WaitGroups vs Other Synchronization Primitives](#-waitgroups-vs-other-synchronization-primitives)
---
## 📚 Introduction to WaitGroups

###  What is a WaitGroup?

> 📝 **_A WaitGroup in Go is a synchronization mechanism from the sync package that allows you to wait for
multiple goroutines to complete their execution before moving forward in your main program._**

👉 It is like a counter-based barrier that blocks until all goroutines finish their work.

Think of it as a counter:

- You increase the counter when you start goroutines.

- Each goroutine decreases the counter when it’s done.

- The main program waits until the counter becomes zero.

### 🎯 Why do we need WaitGroups?

In concurrent programs, multiple goroutines execute tasks ***asynchronously***.

> ⚠️ **WARNING**: Without coordination, your main goroutine may exit before others complete — leading to ***partial execution*** or ***unexpected behavior***.

✅ **WaitGroups** solve this by ensuring that the main goroutine waits until all launched goroutines are done.

### 🔧 The sync.WaitGroup Type

The `sync` package provides the **WaitGroup** type, which has ***three main methods***:

| Method           | Description                                    |
|------------------|------------------------------------------------|
| `Add(delta int)` | Increases (or decreases) the internal counter. |
| `Done()`         | Decrements the counter (same as `Add(-1)`).    |
| `Wait()`         | Blocks the caller until the counter becomes 0. |

---

## 🚀 Basic Usage of WaitGroups

### 📋 How this works step by step:

| Step | Action              | Description                                     |
|------|---------------------|-------------------------------------------------|
| 1️⃣  | `wg.Add(1)`         | Increments counter for each new goroutine       |
| 2️⃣  | `go worker(i, &wg)` | Launches a new concurrent task                  |
| 3️⃣  | `defer wg.Done()`   | Marks this goroutine as done when completed     |
| 4️⃣  | `wg.Wait()`         | Main goroutine waits until counter = 0          |
| 5️⃣  | All Done ✅          | Main goroutine resumes and program exits safely |

---

## ⚙️ Internal Working of WaitGroups

The Go source code defines a simplified version of **WaitGroup** as:
```go
type WaitGroup struct {
    noCopy noCopy    // prevents copying of WaitGroup
    state1 uint64    // stores counter and waiter info
    sema   uint32    // semaphore used for blocking
}
```

### Let’s break this down:

- Counter (state1 high bits) – how many goroutines are being waited on.

- Waiter count (state1 low bits) – how many goroutines are currently waiting (blocked on `Wait()`).

- Semaphore (sema) – used to block and wake up goroutines efficiently.

### 🔄 Lifecycle of a WaitGroup

#### 🆕 Initialization

> 💡 **Good To Know**: A WaitGroup is just a struct, ***no need for explicit init*** (value starts at zero).

#### ➕ 1. Adding (Add(n))

***Increases*** the counter by `n`.  
⚠️ **WARNING**: Must be called ***before*** starting the goroutine.

#### ⏳ 2. Waiting (Wait())

- If counter > 0 → ***blocks*** current goroutine.  
- Uses semaphore internally to suspend goroutines instead of ***busy waiting***.

#### ✅ 3. Done (Done())

***Decreases*** counter by 1.  
If counter == 0 → ***wakes up*** all goroutines blocked in `Wait()`.

### 🤖 How Wait() Works

When you call `wg.Wait()`, the goroutine checks if the counter is zero.

- If ***zero*** → it continues immediately.
- If ***non-zero*** → it sleeps (blocked on semaphore).

When the last goroutine calls `Done()`, the counter becomes zero → ***all blocked goroutines*** waiting on semaphore are released.

---

## 🔑 Key Points

| 🚨 Rule                 | Description                                                                                                          |
|-------------------------|----------------------------------------------------------------------------------------------------------------------|
| ➕ **Add Before Launch** | Always call `wg.Add(1)` ***before*** starting a goroutine                                                            |
| ✅ **Always Done**       | Always use `wg.Done()` inside the goroutine (best practice: `defer wg.Done()`)                                       |
| ⏱️ **Wait Last**        | `Wait()` should only be called ***after*** launching all goroutines                                                  |
| ⚠️ **Balance Counts**   | If `Add()` and `Done()` calls don't match → program will ***hang*** (Wait forever) or ***panic*** (negative counter) |

---

## ✅ Best Practices

### 🎯 Always match Add() with Done()

Every `Add(1)` must have a corresponding `Done()` to avoid ***deadlocks*** or ***panics***.

> 📝 **NOTE**: Call `defer wg.Done()` at the start of the goroutine.

``` go
wg.Add(1)
go func() {
    defer wg.Done()
    // do work
}()
```

### Call Add() before starting the goroutine

Always call `wg.Add(1)` before launching the goroutine to ensure the counter is accurate.

Don’t risk a race condition by calling `Add()` inside the goroutine.

### Don’t call Add() after Wait()

Once `Wait()` starts, adding more goroutines can cause undefined behavior.

Rule: All `Add()` calls must be done before calling `Wait()`.

### 💥 Avoid Negative Counter (Panic)

>⚠️ **WARNING**: If `Done()` is called more times than `Add()`, it ***panics***

### Don’t Copy WaitGroups

A WaitGroup is a `struct`, but copying it leads to race conditions.

> 📝 **NOTE**: **Always pass a pointer `(*sync.WaitGroup)` to goroutines.**

### 📡 Use Channels for Data, WaitGroups for Sync

**WaitGroup** is only for ***synchronization*** (waiting).

> 💡 **Good To Know**: If you also need to pass results back, combine **WaitGroups** with ***channels***.

---

## ⚖️ WaitGroups vs Other Synchronization Primitives

Go provides several synchronization primitives, each serving different purposes.
Let’s see how **WaitGroup** compares with **Channels**, **Mutex**, and **Context**.

### 1️⃣ WaitGroup vs Channel

| Aspect              | WaitGroup                                      | Channel                                        |
|---------------------|------------------------------------------------|------------------------------------------------|
| Purpose             | Synchronize completion of multiple goroutines  | Communicate and share data between goroutines  |
| Communication       | No data transfer                               | Used for sending and receiving values          |
| Usage pattern       | Add → Done → Wait                              | Send → Receive                                 |
| Direction           | One-way signaling (counter-based)              | Two-way data flow                              |
| Typical use case    | Wait for all goroutines to finish              | Pipeline pattern, data passing, worker pools   |
| Internal mechanism  | Counter + semaphore                            | FIFO queue                                     |
| Example use         | Wait until all workers are done                | Pass computed results between goroutines       |

### 2️⃣ WaitGroup vs Mutex

| Aspect           | WaitGroup                         | Mutex                                      |
|------------------|-----------------------------------|--------------------------------------------|
| Purpose          | Wait for goroutines to finish     | Protect shared data from concurrent access |
| Shared state     | None                              | Yes (protects critical section)            |
| Communication    | None                              | None                                       |
| Typical use case | Synchronizing end of work         | Synchronizing access to data               |
| Blocking         | `Wait()` blocks until counter = 0 | Lock() blocks until unlocked               |
| Internal working | Counter-based semaphore           | Lock with ownership control                |

### 3️⃣ WaitGroup vs Context

| Aspect             | WaitGroup                                 | Context                                         |
|--------------------|-------------------------------------------|-------------------------------------------------|
| Purpose            | Wait for goroutines to complete           | Manage lifecycle and cancellation of goroutines |
| Communication      | No cancellation or timeout                | Can cancel or timeout goroutines                |
| Typical use case   | Ensure all goroutines finish before exit  | Gracefully stop goroutines on timeout/shutdown  |
| Internal mechanism | Counter                                   | Deadline/cancel signaling tree                  |
| Example            | Waiting for completion                    | Cancelling ongoing work                         |

### 📊 Summary

 💡 **Good To Know**: Here's a quick comparison of all synchronization primitives:

| Primitive | Purpose                          | Communication | Handles Cancellation | Protects Data |
|-----------|----------------------------------|---------------|----------------------|---------------|
| WaitGroup | Waiting for goroutines to finish | ❌             | ❌                    | ❌             |
| Channel   | Data exchange between goroutines | ✅             | ❌                    | ❌             |
| Mutex     | Protect shared resources         | ❌             | ❌                    | ✅             |
| Context   | Cancel / timeout goroutines      | ❌             | ✅                    | ❌             |
