package timer

const (
	startInTime   int = 5 * 60
	sessionNumber int = 6
	focusTime     int = 25 * 60
	breakTime     int = 5 * 60
	timeLimit         = 59 * 60
	Tick              = 60
)

type ActivityState int

const (
	Init ActivityState = iota
	StartingInState
	FocusState
	BreakState
	TimeoutState
)

type Timer struct {
	Count         int
	TotalCount    int
	StartInTime   int
	SessionNumber int
	FocusTime     int
	BreakTime     int
	StreamTime    int
	IsRunning     bool
	Activity      ActivityState
}

func NewTimer() (timer Timer) {
	timer = Timer{
		StartInTime:   startInTime,
		SessionNumber: sessionNumber,
		FocusTime:     focusTime,
		BreakTime:     breakTime,
		Activity:      Init,
	}
	streamTime := (sessionNumber * (focusTime + breakTime)) + startInTime
	timer.StreamTime = streamTime
	return
}

func (t *Timer) Update() error {
	if t.IsRunning {
		if t.Count == 0 {
			t.changeTimerState()
		}
		t.Count -= 1
		t.TotalCount -= 1
	}
	return nil
}

func (t *Timer) IncreaseStart() {
	t.StartInTime = setTime(t.StartInTime + 60)
	t.calcStreamTime()
}

func (t *Timer) DecreaseStart() {
	t.StartInTime = setTime(t.StartInTime - 60)
	t.calcStreamTime()
}

func (t *Timer) IncreaseSession() {
	t.SessionNumber = limitSession(t.SessionNumber + 1)
	t.calcStreamTime()
}

func (t *Timer) DecreaseSession() {
	t.SessionNumber = limitSession(t.SessionNumber - 1)
	t.calcStreamTime()
}

func (t *Timer) IncreaseFocus() {
	t.FocusTime = setTime(t.FocusTime + 60)
	t.calcStreamTime()
}

func (t *Timer) DecreaseFocus() {
	t.FocusTime = setTime(t.FocusTime - 60)
	t.calcStreamTime()
}

func (t *Timer) IncreaseBreak() {
	t.BreakTime = setTime(t.BreakTime + 60)
	t.calcStreamTime()
}

func (t *Timer) DecreaseBreak() {
	t.BreakTime = setTime(t.BreakTime - 60)
	t.calcStreamTime()
}

func (t *Timer) HandleStart() {
	t.IsRunning = true
	t.Count = t.StartInTime * Tick
	t.TotalCount = t.StreamTime * Tick
}

func (t *Timer) changeTimerState() {
	switch t.Activity {
	case Init:
		return
	case StartingInState:
		t.Activity = FocusState
		t.Count = focusTime * Tick
	case FocusState:
		t.Activity = BreakState
		t.Count = breakTime * Tick
	case BreakState:
		if t.SessionNumber > 0 {
			t.Activity = FocusState
			t.Count = focusTime * Tick
			t.SessionNumber -= 1
		} else {
			t.Activity = TimeoutState
		}
	case TimeoutState:
		return
	}
}

func (t *Timer) calcStreamTime() {
	t.StreamTime = (t.SessionNumber * (t.FocusTime + t.BreakTime)) + t.StartInTime
}

func setTime(time int) int {
	if time < 60 {
		return 60
	} else if time > timeLimit {
		return timeLimit
	} else {
		return time
	}
}

func limitSession(sessionNumber int) int {
	if sessionNumber < 1 {
		return 1
	} else if sessionNumber > 16 {
		return 16
	} else {
		return sessionNumber
	}
}
