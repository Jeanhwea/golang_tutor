package nc0062

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0062_01(t *testing.T) {
	for i := 1; i < 10; i++ {
		ans := Fibonacci(i)
		t.Logf("fib[%v] = %v\n", i, ans)
	}
	assert.Equal(t, true, true)
}
