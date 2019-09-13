package task

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

type Job interface {
	Do()
}

func New(ctx context.Context, done chan struct{}, max int) *Task {
	task := &Task{
		ctx:         ctx,
		maxInFlight: max,
		done:        done,
		controlChan: make(chan struct{}, max),
		jobChan:     make(chan Job, 10000),
		running:     true,
	}
	go func() {
		// 正常时干活
		task.run()
		// 收到退出信号时，处理剩余未完的任务
		fmt.Printf("handling remaining %d jobs\n", len(task.jobChan))
		for job := range task.jobChan {
			job.Do()
		}
		// 佯装等待几秒，再退出
		for i := 3; i >= 1; i-- {
			fmt.Printf("shutdown in %d second\n", i)
			time.Sleep(time.Second)
		}
		task.done <- struct{}{}
	}()
	return task
}

type Task struct {
	running     bool
	jobChan     chan Job
	done        chan struct{}
	controlChan chan struct{}
	maxInFlight int
	ctx         context.Context
}

func (t *Task) run() {
	for {
		select {

		case job := <-t.jobChan:
			t.controlChan <- struct{}{}
			go func() {
				job.Do()
				<-t.controlChan
			}()

		case <-t.ctx.Done():

			fmt.Println("will stop")
			go t.stop()

			fmt.Println("prepare to shutdown...")
			return
		}
	}
}

// 停止处理任务，关闭chan
func (t *Task) stop() {
	mu.Lock()
	defer mu.Unlock()
	t.running = false
	close(t.jobChan)
	fmt.Println("stopped")
}

func (t *Task) Push(job Job) bool {
	mu.Lock()
	defer mu.Unlock()

	if !t.running {
		return false
	}

	t.jobChan <- job
	return true
}
