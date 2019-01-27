package taskrunner

import (
	"time"
)

type Worker struct {
	ticker *time.Ticker //不断接受系统发过来的时间，达到一定的时间间隔的时候，就自动触发我们想要做的事情
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
		select { //用这种方式不会阻塞
		case <- w.ticker.C: //接收 通过channel形式传过来的 NewTicker(interval * time.Second),里面的时间，进而再执行任务
			go w.runner.StartAll()
		}
	}
}




func Start() {
	r := NewRunner(3,true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r) //每3秒执行 依次Runner
	go w.startWorker()
}

//整体步骤

// user -> api service -> delete video
// api service -> scheduler -> write video deletion record
// timer
// timer -> runner -> read wvdr -> exec -> delete video from folder