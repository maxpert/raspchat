package rascore

import (
    "log"
    "time"
)

type StopWatch struct {
    start, stop time.Time
    name        string
}

func StartStopWatch(name string) *StopWatch {
    return &StopWatch{
        name:  name,
        start: time.Now(),
    }
}

func (self *StopWatch) milliseconds() uint32 {
    return uint32(self.stop.Sub(self.start) / time.Millisecond)
}

func (self *StopWatch) Stop() uint32 {
    self.stop = time.Now()
    return self.milliseconds()
}

func (self *StopWatch) LogDuration() {
    log.Println("Time taken by", self.name, "=", self.Stop(), "ms")
}
