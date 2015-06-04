package my

import (
	"time"
)

type Interval int

const (
	Minutely Interval = iota - 3
	Hourly
	Daily
)

func (i Interval) Do(f func(time.Time, *time.Ticker)) {
	now := time.Now()
	var tofuture time.Duration
	var dur time.Duration
	switch i {
	case Minutely:
		dur = 1 * time.Minute
		tofuture = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()+1, 0, 0, time.Local).Sub(now)
	case Hourly:
		dur = 1 * time.Hour
		tofuture = time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, time.Local).Sub(now)
	case Daily:
		dur = 24 * time.Hour
		tofuture = time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local).Sub(now)
	default:
		dur = time.Duration(i) * time.Hour
		tofuture = time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+int(i), 0, 0, 0, time.Local).Sub(now)
	}
	tick := time.NewTicker(dur)
	f(time.Now(), tick)
	go func() {
		<-time.After(tofuture)
		for _ = range tick.C {
			f(time.Now(), tick)
		}
	}()
	return
}
