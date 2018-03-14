package isp

import "time"

type (
	Ticker interface {
		Start()
		Wait() time.Time
		Stop()
	}
	IntervalTicker struct {
		interval time.Duration
		ticker   *time.Ticker
	}
	LimitlessTicker struct{}
)

func NewTicker(interval time.Duration) Ticker {
	if interval == 0{
		return NewLimitlessTicker()
	}
	return NewIntervalTicker(interval)
}


func NewIntervalTicker(interval time.Duration) *IntervalTicker {
	return &IntervalTicker{
		interval:interval,
	}
}
func (s *IntervalTicker) Start() {
	s.ticker = time.NewTicker(s.interval)
}
func (s *IntervalTicker) Wait() time.Time {
	return <-s.ticker.C
}
func (s *IntervalTicker) Stop() {
	s.ticker.Stop()
	s.ticker = nil
}
func (s *IntervalTicker) SetInterval(interval time.Duration) {
	s.interval = interval
}
func (s *IntervalTicker) GetInterval() (interval time.Duration) {
	return s.interval
}

func NewLimitlessTicker() *LimitlessTicker {
	return &LimitlessTicker{}
}
func (s *LimitlessTicker) Start() {
}
func (s *LimitlessTicker) Wait() time.Time {
	return time.Now()
}
func (s *LimitlessTicker) Stop() {
}
