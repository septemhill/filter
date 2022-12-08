package task

type Task interface {
	Do()
}

type TaskFunc func()

func (f TaskFunc) Do() {
	f()
}
