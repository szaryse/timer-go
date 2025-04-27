package timer

type Timer struct {
	Count int
}

func NewTimer() Timer {
	return Timer{}
}

func (t *Timer) Update() error {
	t.Count += 1
	return nil
}
