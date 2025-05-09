package retry

import (
	"math"
	"time"
)

// Retry retries the given function until no error is returned or max retries are reached.
// f is the function to be retried; it must return an error on failure or nil on success.
// expRate defines the exponential backoff base duration between retry attempts.
// maxRetries specifies the maximum number of retry attempts allowed.
//
// The function implements exponential backoff algorithm for retry delays:
// t = b^c
// where:
//   - t is the time delay between retry attempts (in seconds)
//   - b is the base duration (expRate in seconds)
//   - c is the retry attempt number (0-based)
//
// Source: https://en.wikipedia.org/wiki/Exponential_backoff
//
// For example, with expRate=2s and maxRetries=4, the delays will be:
//   - 1st retry: 2^0 = 1s
//   - 2nd retry: 2^1 = 2s
//   - 3rd retry: 2^2 = 4s
//   - 4th retry: 2^3 = 8s
//
// Returns the last encountered error if all retries fail, or nil if the function eventually succeeds.
func Retry(f func() error, expRate time.Duration, maxRetries int) error {
	var err error
	b := expRate.Seconds()
	for i := 0; i < maxRetries; i++ {
		err = f()
		if err == nil {
			return nil
		}

		// Exponential backoff algorithm
		// https://en.wikipedia.org/wiki/Exponential_backoff
		// t = b^c
		// where t is the time delay applied between actions
		// and b is multiplicative factor or base
		// and c is the number of adverse events observed, the value of c is incremented each time an adverse event is observed
		t := math.Pow(b, float64(i))
		sleep := time.Duration(t * float64(time.Second))
		time.Sleep(sleep)
	}
	return err
}
