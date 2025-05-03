package timer

const (
	startInTime   int = 5 * 60
	sessionNumber int = 6
	focusTime     int = 25 * 60
	breakTime     int = 5 * 60
	timeLimit         = 59 * 60
	Tick              = 60
)

type Timer struct {
	Count         int
	StartInTime   int
	SessionNumber int
	FocusTime     int
	BreakTime     int
	StreamTime    int
	IsRunning     bool
}

func NewTimer() (timer Timer) {
	timer = Timer{
		StartInTime:   startInTime,
		SessionNumber: sessionNumber,
		FocusTime:     focusTime,
		BreakTime:     breakTime,
	}
	streamTime := (sessionNumber * (focusTime + breakTime)) + startInTime
	timer.Count = streamTime * Tick
	timer.StreamTime = streamTime
	return
}

func (t *Timer) Update() error {
	if t.IsRunning {
		t.Count -= 1
	}
	return nil
}

func (t *Timer) HandleAction(action string) {
	switch action {
	case "increaseStart":
		t.StartInTime = setTime(t.StartInTime + 60)
	case "decreaseStart":
		t.StartInTime = setTime(t.StartInTime - 60)
	case "increaseSession":
		t.SessionNumber = limitSession(t.SessionNumber + 1)
	case "decreaseSession":
		t.SessionNumber = limitSession(t.SessionNumber - 1)
	case "increaseFocus":
		t.FocusTime = setTime(t.FocusTime + 60)
	case "decreaseFocus":
		t.FocusTime = setTime(t.FocusTime - 60)
	case "increaseBreak":
		t.BreakTime = setTime(t.BreakTime + 60)
	case "decreaseBreak":
		t.BreakTime = setTime(t.BreakTime - 60)
	case "start":
		t.IsRunning = true
	}
	t.calcStreamTime()
}

func (t *Timer) calcStreamTime() {
	t.StreamTime = (t.SessionNumber * (t.FocusTime + t.BreakTime)) + t.StartInTime
	t.Count = t.StreamTime * Tick
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
