# MSCS-632-M51_Assignment6

## Features

- Concurrent task processing using multiple workers
- Thread-safe task queue implementation
- Result collection with proper synchronization
- Logging of worker activities
- File output for results (Java implementation)

## Running the Code

### Go Implementation

1. Navigate to the Go directory:
```bash
cd Go
```

2. Run the program:
```bash
go run main.go
```

The Go implementation will:
- Create 3 worker goroutines
- Process 10 tasks
- Print results to console

### Java Implementation

1. Navigate to the Java directory:
```bash
cd Java
```

2. Compile the Java files:
```bash
javac *.java
```

3. Run the program:
```bash
java Main
```

The Java implementation will:
- Create 4 worker threads
- Process 10 tasks
- Save results to `results.txt`
- Log activities to console

## Implementation Details

### Go Version
- Uses channels for task distribution
- Implements mutex for result synchronization
- Uses WaitGroup for worker coordination

### Java Version
- Uses a custom thread-safe TaskQueue
- Implements synchronized blocks for result collection
- Uses Thread.join() for worker coordination
- Includes file output and logging

## Expected Output

Both implementations will process 10 tasks concurrently and produce similar output, though the Java version will save results to a file while the Go version prints to console.

## Requirements

- Go 1.x or later
- Java 8 or later