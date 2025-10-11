# 🚀 Go Concurrency Quizzes

A comprehensive collection of practical Go concurrency exercises and coding challenges designed to help you master goroutines, channels, context, and advanced concurrency patterns.

## 📚 Overview

This project contains a series of hands-on coding tasks that cover essential Go concurrency concepts. Each quiz focuses on a specific concurrency pattern or technique, providing you with practical experience in writing concurrent Go code.

## 🎯 What You'll Learn

- **Goroutines**: Lightweight threads and concurrent execution
- **Channels**: Communication and synchronization between goroutines
- **Context**: Cancellation, timeouts, and request-scoped values
- **Worker Pools**: Managing concurrent workloads efficiently
- **Pipeline Patterns**: Building data processing pipelines
- **Error Handling**: Managing errors in concurrent operations
- **Stream Processing**: Working with streams of data
- **Advanced Patterns**: Fan-out/fan-in, tee, bridge, and more

## 📁 Available Quizzes

### 🔄 Basic Channel Operations

- **[Channel Relay with Context Cancellation](Channel%20Relay%20with%20Context%20Cancellation/)** - Implement `orDone` pattern with context cancellation
- **[Merging Two Sorted Channels](Merging%20Two%20Sorted%20Channels/)** - Merge sorted data from multiple channels
- **[Merging Channels Concurrently](Merging%20Channels%20Concurrently/)** - Combine multiple channels into one
- **[Channel Chaining with inc](Channel%20Chaining%20with%20inc/)** - Create chains of channel transformations

### 🏗️ Worker Patterns

- **[Worker Pool with Goroutines and Channels](Worker%20Pool%20with%20Goroutines%20and%20Channels/)** - Implement a simple worker pool
- **[Controlled Parallel Execution with Error Aggregation](Controlled%20Parallel%20Execution%20with%20Error%20Aggregation/)** - Manage parallel tasks with error handling
- **[Concurrent Function Execution with Error Propagation](Concurrent%20Function%20Execution%20with%20Error%20Propagation/)** - Execute functions concurrently and propagate any errors
- **[Concurrent Task Execution with Semaphore Limiting](Concurrent%20Task%20Execution%20with%20Semaphore%20Limiting/)** - Control concurrent execution using semaphore pattern

### 🔍 Search and Processing

- **[Concurrent Search with Early Results and Context Cancellation](Concurrent%20Search%20with%20Early%20Results%20and%20Context%20Cancellation/)** - Search across multiple sources with early termination
- **[Concurrent Sorted Head from Multiple Readers](Concurrent%20Sorted%20Head%20from%20Multiple%20Readers/)** - Get sorted results from concurrent readers

### 📥 I/O and Downloads

- **[Concurrent Downloads with Error Aggregation](Concurrent%20Downloads%20with%20Error%20Aggregation/)** - Download multiple files concurrently with error handling
- **[Running a Task with Context Timeout](Running%20a%20Task%20with%20Context%20Timeout/)** - Execute tasks with timeout control

### 🔄 Stream Processing

- **[Context-Aware Generator and Squarer Pipeline](Context-Aware%20Generator%20and%20Squarer%20Pipeline/)** - Build context-aware data pipelines
- **[repeatFn and take with Context Cancellation](repeatFn%20and%20take%20with%20Context%20Cancellation/)** - Implement functional stream operations
- **[Duplicating a Stream with tee](Duplicating%20a%20Stream%20with%20tee/)** - Split a stream into multiple outputs
- **[Flattening Streams of Streams with bridge](Flattening%20Streams%20of%20Streams%20with%20bridge/)** - Flatten nested stream structures
- **[Ring Buffer with Overwriting Writes](Ring%20Buffer%20with%20Overwriting%20Writes/)** - Implement a circular buffer that overwrites old data when full
- **[Coordinated Producer Shutdown with Context and WaitGroup](Coordinated%20Producer%20Shutdown%20with%20Context%20and%20WaitGroup/)** - Coordinate multiple producers with graceful shutdown

### 🔧 Synchronization Primitives

- **[Custom WaitGroup with Channel Synchronization](Custom%20WaitGroup%20with%20Channel%20Synchronization/)** - Build a custom WaitGroup using only channels
- **[Channel-Based Implementation of <once>](Channel-Based%20Implementation%20of%20%3Conce%3E/)** - Implement sync.Once behavior using channels

## 🚀 Getting Started

### Prerequisites

- Go 1.19 or later
- Basic understanding of Go syntax and concepts

### Setup

1. Clone this repository:

   ```bash
   git clone <repository-url>
   cd go-concurrency-quizzes
   ```

2. Navigate to any quiz directory:

   ```bash
   cd "Channel Relay with Context Cancellation"
   ```

3. Read the README.md file to understand the task requirements

4. Implement the required functions in `main.go`

5. Run your solution:
   ```bash
   go run main.go
   ```

## 🎯 How to Use

1. **Choose a Quiz**: Start with basic concepts and work your way up to advanced patterns
2. **Read the Requirements**: Each quiz has a detailed README explaining what to implement
3. **Implement the Solution**: Write your code in the provided `main.go` file
4. **Test Your Code**: Run the program to verify your implementation works correctly
5. **Study the Patterns**: Understand how the concurrency patterns work and when to use them

## 📖 Learning Path

### Beginner Level

- Channel Relay with Context Cancellation
- Merging Two Sorted Channels
- Running a Task with Context Timeout
- Channel Chaining with inc

### Intermediate Level

- Worker Pool with Goroutines and Channels
- Concurrent Downloads with Error Aggregation
- Merging Channels Concurrently
- Concurrent Function Execution with Error Propagation
- Ring Buffer with Overwriting Writes
- Concurrent Task Execution with Semaphore Limiting
- Custom WaitGroup with Channel Synchronization

### Advanced Level

- Concurrent Search with Early Results
- Stream Processing Patterns (tee, bridge)
- Context-Aware Pipelines
- Coordinated Producer Shutdown with Context and WaitGroup
- Channel-Based Implementation of <once>

## 🔧 Tips for Success

- **Start Simple**: Begin with basic channel operations before tackling complex patterns
- **Understand Context**: Context is crucial for cancellation and timeout scenarios
- **Handle Errors**: Always consider error handling in concurrent operations
- **Avoid Goroutine Leaks**: Ensure goroutines terminate properly
- **Use Buffered Channels**: When appropriate, buffered channels can improve performance
- **Test Thoroughly**: Concurrent code can have race conditions - test with various scenarios

## 🤝 Contributing

Feel free to:

- Add new concurrency quizzes
- Improve existing exercises
- Fix bugs or clarify instructions
- Add more detailed explanations

## 📚 Additional Resources

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Go by Example: Goroutines](https://gobyexample.com/goroutines)
- [Go by Example: Channels](https://gobyexample.com/channels)
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

**Happy coding! 🎉** Master these concurrency patterns and you'll be well-equipped to build robust, scalable Go applications.
