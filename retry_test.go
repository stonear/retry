package retry

import (
	"errors"
	"log"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	testCounter := 0

	testCases := []struct {
		retryFunc  func() error
		expRate    time.Duration
		maxRetries int
		expect     error
	}{
		{
			retryFunc: func() error {
				// empty func
				log.Printf("retrying testcase 1 at %v", time.Now())
				return nil
			},
			expRate:    time.Second,
			maxRetries: 5,
			expect:     nil,
		},
		{
			retryFunc: func() error {
				// error func
				log.Printf("retrying testcase 2 at %v", time.Now())
				return errors.New("expected error")
			},
			expRate:    time.Duration(0.75 * float64(time.Second)),
			maxRetries: 3,
			expect:     errors.New("expected error"),
		},
		{
			retryFunc: func() error {
				// do something
				log.Printf("retrying testcase 3 at %v", time.Now())
				testCounter++
				if testCounter < 3 {
					return errors.New("expected error")
				}
				return nil
			},
			expRate:    time.Duration(1.5 * float64(time.Second)),
			maxRetries: 3,
			expect:     nil,
		},
	}

	for _, tc := range testCases {
		err := Retry(tc.retryFunc, tc.expRate, tc.maxRetries)
		assert.Equal(t, tc.expect, err)
	}
}

func TestAlg(t *testing.T) {
	testCases := []struct {
		b      time.Duration
		c      int
		expect time.Duration
	}{
		{
			b:      time.Second,
			c:      5,
			expect: time.Duration(5.0 * float64(time.Second)),
		},
		{
			b:      time.Duration(0.75 * float64(time.Second)),
			c:      3,
			expect: time.Duration(2.3125 * float64(time.Second)),
		},
		{
			b:      time.Duration(1.5 * float64(time.Second)),
			c:      3,
			expect: time.Duration(4.75 * float64(time.Second)),
		},
		{
			// An exponential backoff algorithm where b = 2 is referred to as a binary exponential backoff algorithm.
			b:      time.Duration(2.0 * float64(time.Second)),
			c:      3,
			expect: time.Duration(7 * float64(time.Second)),
		},
	}
	for _, tc := range testCases {
		var totalElapsed time.Duration
		for i := 0; i < tc.c; i++ {
			t := math.Pow(tc.b.Seconds(), float64(i))
			totalElapsed += time.Duration(t * float64(time.Second))
		}
		assert.Equal(t, tc.expect, totalElapsed)
	}
}
