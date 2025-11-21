# In-Memory Cache (Go)

A simple, thread-safe in-memory cache for Go.  
Stores values in RAM with optional TTL (time-to-live) expiration.

---

## Features

- `Set` / `Get` / `Delete` / `Clear`
- TTL expiration support
- Thread-safe (`sync.RWMutex`)
- Typed errors (`ErrNotFound`, `ErrExpired`, `ErrInvalidTTL`)

---

## Installation

```bash
go get github.com/brogrammer17/cache
```

## Usage

```go
package main

import (
"fmt"
"time"

    "github.com/brogrammer17/cache"
)

func main() {
    // Create a new cache
    c := cache.New()

    // Set a value with TTL 5 seconds
    err := c.Set("username", "Damir", 5*time.Second)
    if err != nil {
        fmt.Println("Set error:", err)
        return
    }

    // Get the value
    value, err := c.Get("username")
    if err != nil {
        fmt.Println("Get error:", err)
    } else {
        fmt.Println("Value:", value)
    }

    // Wait for 6 seconds so the value expires
    time.Sleep(6 * time.Second)

    value, err = c.Get("username")
    if err != nil {
        switch err {
        case cache.ErrNotFound:
            fmt.Println("Key not found")
        case cache.ErrExpired:
            fmt.Println("Key expired")
        default:
            fmt.Println("Unknown error:", err)
        }
    } else {
        fmt.Println("Value:", value)
    }

    // Delete the key
    c.Delete("username")

    // Clear the cache
    c.Clear()
}
```

## License

MIT License © 2025 Damir Hasanshin
