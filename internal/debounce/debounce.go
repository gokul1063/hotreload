package debounce

import (
	"time"
)

type Debouncer struct {
	delay time.Duration
	timer *time.Timer
	out   chan struct{}
}

func New(delay time.Duration) *Debouncer {

	return &Debouncer{
		delay: delay,
		out:   make(chan struct{}, 1),
	}
}

func (d *Debouncer) Trigger() {

	if d.timer != nil {
		d.timer.Stop()
	}

	d.timer = time.AfterFunc(d.delay, func() {

		select {
		case d.out <- struct{}{}:
		default:
		}

	})
}

func (d *Debouncer) Events() <-chan struct{} {
	return d.out
}
