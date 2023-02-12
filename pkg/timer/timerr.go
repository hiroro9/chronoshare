package timer

import (
	"time"
)

type Timer struct {
	id         string // unique identifier
	initRemain int
	remain     int
	stop       bool
}

func (t *Timer) Start() {

	for {
		if !t.stop {
			t.remain -= 1
		}
		time.Sleep(1000000000)
	}
}

func (t *Timer) Stop() {
	t.stop = true
}

func (t *Timer) Restart() {
	t.stop = false
}

func (t *Timer) Reset() {
	t.remain = t.initRemain
}

func (t *Timer) GetId() string {
	return t.id
}

func (t *Timer) SetId(id string) {
	t.id = id
}

func (t *Timer) GetRemain() int {
	return t.remain
}

func (t *Timer) SetRemain(remain int) {
	t.remain = remain
}

func NewTimer(id string, remain int) *Timer {
	return &Timer{
		id:         id,
		initRemain: remain,
		remain:     remain,
		stop:       false,
	}
}
