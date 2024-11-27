# Go-Cache

A lightweight, thread-safe in-memory cache implementation in Go with TTL support.

## Features

- Simple key-value storage
- Thread-safe operations with mutex locks
- TTL (Time-To-Live) support for cached items
- Basic operations: Set, Get, Delete
- Count functionality for cached items
- TCP server implementation for network access

## Installation

```bash
git clone https://github.com/yourusername/go-cache.git
cd go-cache
go build
```

## Usage

### As a Package

```go
import "your-module/db"

// Create a new database instance
cache := db.NewDB()

// Set a value with TTL
cache.Set("key", "value", 3600) // TTL in seconds

// Get a value
value, err := cache.Get("key")
if err != nil {
    // Handle error (not found or expired)
}

// Delete a value
cache.Delete("key")

// Get count of active items
count := cache.Count()
```

### As a TCP Server

The project includes a TCP server implementation that listens on port 3169.

To run the server:

```bash
go run main.go
```

## Project Structure

```
.
├── main.go         # TCP server implementation
├── db
│   └── db.go      # Core cache implementation
└── README.md
```

## Implementation Details

### Cache Operations

- `Set(key, value string, ttlSeconds int64)`: Stores a value with an optional TTL
- `Get(key string) (string, error)`: Retrieves a value, returns error if not found or expired
- `Delete(key string)`: Removes a value from the cache
- `Count() int`: Returns the count of non-expired items in the cache

### Concurrency

The implementation uses `sync.RWMutex` to ensure thread-safe operations on the cache.