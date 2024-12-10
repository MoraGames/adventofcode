package utils

import "time"

type Timer struct {
	Start time.Time
	Inter []time.Time
	End   time.Time
}

func NewTimer(alreadyStarted bool) *Timer {
	if alreadyStarted {
		return &Timer{Start: time.Now()}
	}
	return &Timer{}
}

func (t *Timer) StartTimer() {
	t.Start = time.Now()
}

func (t *Timer) InterTimer() {
	t.Inter = append(t.Inter, time.Now())
}

func (t *Timer) EndTimer() {
	t.End = time.Now()
}

func (t *Timer) ElapsedTime() time.Duration {
	return t.End.Sub(t.Start)
}

func (t *Timer) InterTimes() []time.Duration {
	interTimes := make([]time.Duration, 0, len(t.Inter))
	for _, interTime := range t.Inter {
		interTimes = append(interTimes, interTime.Sub(t.Start))
	}
	return interTimes
}

func (t *Timer) DeltaTimes() []time.Duration {
	deltaTimes := make([]time.Duration, 0, len(t.Inter))
	for i := 0; i < len(t.Inter); i++ {
		if i == 0 {
			deltaTimes = append(deltaTimes, t.Inter[i].Sub(t.Start))
			continue
		}
		if i == len(t.Inter)-1 {
			deltaTimes = append(deltaTimes, t.End.Sub(t.Inter[i]))
			continue
		}
		deltaTimes = append(deltaTimes, t.Inter[i+1].Sub(t.Inter[i]))
	}
	return deltaTimes
}
