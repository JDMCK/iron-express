package core

type Timer struct {
	Cycles     int
	elapsed    int
	duration   int
	maxCycles  int // if repeat is true, once max cycles is reached, the timer restarts from 0
	repeat     bool
	paused     bool
	onComplete func()
}

func NewTimer(duration, maxCycles int, repeat bool, onComplete func()) Timer {
	return Timer{
		elapsed:   0,
		duration:  duration,
		maxCycles: maxCycles,
		Cycles:    0,
		repeat:    repeat,
		paused:    false,
	}
}

func (t *Timer) TogglePause() {
	t.paused = !t.paused
}

func (t *Timer) Pause() {
	t.paused = true
}

func (t *Timer) Play() {
	t.paused = false
}

func (t *Timer) Restart() {
	t.Cycles = 0
	t.elapsed = 0
	t.paused = false
}

func (t *Timer) Reset() {
	t.Cycles = 0
	t.elapsed = 0
	t.paused = true
}

func (t *Timer) Update() {
	if t.paused || (!t.repeat && t.Cycles >= t.maxCycles-1) { // -1 because the 0th cycle is the first cycle
		return
	}

	if t.elapsed >= t.duration {
		t.elapsed -= t.duration
		t.Cycles += 1

		if t.Cycles >= t.maxCycles {
			t.Cycles = 0
		}
	}

	t.elapsed++
}
