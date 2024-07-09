package retry

import (
	"math"
	"time"
)

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
		t := math.Pow(b, float64(i+1))
		sleep := time.Duration(t * float64(time.Second))
		time.Sleep(sleep)
	}
	return err
}
