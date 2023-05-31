package nc0100

import (
	"testing"
)

func Test_NC0100_01(t *testing.T) {
	s := Constructor(2)
	s.set(1, 1)
	s.set(2, 2)
	output := s.get(1)
	t.Log(output)
	s.set(3, 3)
	output = s.get(2)
	t.Log(output)
	s.set(4, 4)
	output = s.get(1)
	t.Log(output)
	output = s.get(3)
	t.Log(output)
	output = s.get(4)
	t.Log(output)
}

func Test_NC0100_02(t *testing.T) {
	s := Constructor(1)
	s.set(1, 1)
	s.set(1, 2)
	output := s.get(1)
	t.Log(output)
}
