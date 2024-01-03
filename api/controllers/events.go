package controllers

import (
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var count = 0

func (a *App) startTimerEmit() {
	for {
		timeNow := time.Now().Format("15:04:05")
		runtime.EventsEmit(a.ctx, "app:tick", timeNow)
		time.Sleep(1 * time.Second)	
	}
}

func (a *App) TimerOff() {
	runtime.EventsOff(a.ctx, "tickOn")
}

func (a *App) TimerMultiple() {
	// Count 5 times
	runtime.EventsOnMultiple(a.ctx, "tickMultiple", func(data ...interface{}) {
		count++
		runtime.EventsEmit(a.ctx, "app:tickMultiple", count)
	}, 5)
}

func (a *App) Count() (count int) {
	count++
	return
}
