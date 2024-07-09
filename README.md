# Retry
[![Go Reference](https://pkg.go.dev/badge/github.com/stonear/retry.svg)](https://pkg.go.dev/github.com/stonear/retry)

Very simple Go package to retry a function using exponential backoff algorithm.

## Installation

```
go get github.com/stonear/retry
```


## Usage

```go
import "github.com/stonear/retry"

// Define your function
func yourFunction() error {
    // Your code here
    return nil
}

// Call the Retry function
err := retry.Retry(yourFunction, time.Second, 5)
if err != nil {
    // Handle the error
}
```

# Contributing
Contributions are welcome. Please open a pull request with your changes.
