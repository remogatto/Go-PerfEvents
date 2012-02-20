package perf

import (
	"testing"
)

func TestPerf(t *testing.T) {
	counter1, err := NewCounter_Instructions(true, false)
	if err != nil {
		t.Fatal(err)
	}

	_, err = counter1.Read()
	if err != nil {
		t.Fatal(err)
	}

	err = counter1.Close()
	if err != nil {
		t.Fatal(err)
	}
}
