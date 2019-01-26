package taskrunner

import (
	"time"
)

type Worker struct {
	ticker *time.Ticker //
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker{
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() { // Worker内置方法
	for {
		select {
		case <- w.ticker.C:
			go w.runner.StartAll()
		}
	}
}




func Start() {
	r := NewRunner(3,true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r) //每3秒执行 依次Runner
	go w.startWorker()
}