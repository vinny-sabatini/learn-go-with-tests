package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("validate output for countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		countdownOperationsSpy := &CountdownOperationsSpy{}

		Countdown(buffer, countdownOperationsSpy)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		countdownOperationsSpy := &CountdownOperationsSpy{}
		Countdown(countdownOperationsSpy, countdownOperationsSpy)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, countdownOperationsSpy.Calls) {
			t.Errorf("wanted calls %v, got %v", want, countdownOperationsSpy.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime)
	}
}
