package task

import (
	"errors"
	"sync"
	"time"
)

const (
	defaultTaskBuffer      = 10
	defaultMaxRoutines     = 10
	defaultWorkingRoutines = 4
)

type taskHandler struct {
	taskCh       chan Task
	shrinkCh     chan struct{}
	maxRoutines  uint
	expRoutines  uint
	currRoutines uint
	wg           sync.WaitGroup
	mu           sync.RWMutex
}

func (th *taskHandler) createWorkers(num uint) {
	for i := uint(0); i < num; i++ {
		th.addWorker()
	}
}

func (th *taskHandler) addWorker() {
	th.wg.Add(1)
	th.mu.Lock()
	th.currRoutines += 1
	th.mu.Unlock()

	go func() {
		defer func() {
			th.wg.Done()
			th.mu.Lock()
			th.currRoutines -= 1
			th.mu.Unlock()
		}()

		for {
			select {
			case task := <-th.taskCh:
				task.Do()
			case <-th.shrinkCh:
				return
			}
		}
	}()
}

func (th *taskHandler) Send(task Task) {
	th.taskCh <- task
}

func (th *taskHandler) SendWithTimeout(task Task, timeout time.Duration) error {
	t := time.After(timeout)

	select {
	case <-t:
		return errors.New("task sent timeout")
	case th.taskCh <- task:
	}

	return nil
}

func (th *taskHandler) Shrink(num uint) (uint, error) {
	th.mu.Lock()
	defer th.mu.Unlock()

	prev := th.expRoutines

	if num == 0 {
		return prev, errors.New("num should greater than zero")
	}

	if th.expRoutines-num > 0 {
		th.expRoutines -= num
		go func() {
			for i := uint(0); i < num; i++ {
				th.shrinkCh <- struct{}{}
			}
		}()
		return prev, nil
	}

	return prev, errors.New("over number of working routines")
}

func (th *taskHandler) Grow(num uint) (uint, error) {
	th.mu.Lock()
	defer th.mu.Unlock()

	prev := th.expRoutines

	if num == 0 {
		return prev, errors.New("num should greater than zero")
	}

	if th.expRoutines+num <= th.maxRoutines {
		th.expRoutines += num
		go th.createWorkers(num)
		return prev, nil
	}

	return prev, errors.New("over maximum limitation")
}

func NewTaskHandler(opts ...Option) *taskHandler {
	th := &taskHandler{
		taskCh:      make(chan Task, defaultTaskBuffer),
		shrinkCh:    make(chan struct{}, defaultMaxRoutines),
		maxRoutines: defaultMaxRoutines,
		expRoutines: defaultWorkingRoutines,
	}

	for _, opt := range opts {
		opt.apply(th)
	}

	th.createWorkers(th.expRoutines)

	return th
}
