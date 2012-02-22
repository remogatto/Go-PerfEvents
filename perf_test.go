package perf

import (
	"testing"
)

func Test_NewCounter_Instructions(t *testing.T) {
	counter, err := NewCounter_Instructions(true, false)
	if err != nil {
		t.Fatal(err)
	}

	_, err = counter.Read()
	if err != nil {
		t.Fatal(err)
	}

	err = counter.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_NewEvent(t *testing.T) {
	counter, err := NewCounter(TYPE_SOFTWARE, SW_TASK_CLOCK, 0)
	if err != nil {
		t.Fatal(err)
	}

	_, err = counter.Read()
	if err != nil {
		t.Fatal(err)
	}

	err = counter.Close()
	if err != nil {
		t.Fatal(err)
	}
}
