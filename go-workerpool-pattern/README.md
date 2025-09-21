# Go Worker Pool Example

This project demonstrates two concurrent worker pool patterns in Go:
- **Single-Type Task Worker Pool**: Processes a fixed set of tasks of one type concurrently.
- **Multi-Type Task Worker Pool**: Processes different types of tasks (e.g., email sending, image processing) concurrently using interfaces.

## Project Structure

- `main.go`: Entry point. Contains functions to test both worker pool implementations.
- `workerpool.go`: Implements a worker pool for a single type of task (`Task`).
- `workerpool2.go`: Implements a worker pool for multiple types of tasks using the `MultiTask` interface.
- `go.mod`, `go.sum`: Go module files.

## How It Works

### Single-Type Task Worker Pool
- Creates 20 tasks of type `Task`.
- Processes them concurrently using a pool of 6 workers.

### Multi-Type Task Worker Pool
- Creates a mix of `EmailTask` and `ImageProcessingTask`.
- Processes them concurrently using a pool of 3 workers.

## Running the Project

1. Ensure you have Go installed (version 1.18+ recommended).
2. Clone or download this repository.
3. In `main.go`, comment/uncomment the function calls in `main()` to test either implementation:
```go
WorkerPoolWithOneTypeOfTask()
WorkerPoolWithMultipleTypeOfTasks()
```
4. Run the project:
```sh
go run main.go workerpool.go workerpool2.go
```

## Example Output

```
Processing task with ID: 1
Processing task with ID: 2
... (other tasks)
All tasks completed.
```
Or for multi-type tasks:
```
Sending email to: abc
Processing image from URL: ABC
... (other tasks)
All tasks completed.
```

## License

This project is for educational purposes.
