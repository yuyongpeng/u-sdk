package retryu

import (
	"fmt"
	"math"
	"testing"
)

func TestRetry1(t *testing.T) {
	n := uint(math.Floor(math.Log2(float64(6))))
	fmt.Println(n)
	retry1()
}
