package task

type Option interface {
	apply(*taskHandler)
}

type WithNumberOfWorkers uint

func (opt WithNumberOfWorkers) apply(th *taskHandler) {
	th.expRoutines = uint(opt)
}

type WithMaximumWorkers uint

func (opt WithMaximumWorkers) apply(th *taskHandler) {
	th.maxRoutines = uint(opt)
	th.shrinkCh = make(chan struct{}, uint(opt))
}

type WithMaximumTaskBufferSize uint

func (opt WithMaximumTaskBufferSize) apply(th *taskHandler) {
	th.taskCh = make(chan Task, uint(opt))
}
