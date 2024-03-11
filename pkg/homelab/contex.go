package homelab

type Context struct {
	taskQueue chan QueuedTask
}

type QueuedTask struct {
	serviceID string
	taskID    string
	doneChan  chan bool
	errChan   chan error
}

func NewContext() Context {
	return Context{
		taskQueue: make(chan QueuedTask),
	}
}

func (c *Context) QueueTask(serviceID string, taskID string) (<-chan bool, <-chan error) {
	doneChan := make(chan bool, 1)
	errChan := make(chan error, 1)

	c.taskQueue <- QueuedTask{
		serviceID: serviceID,
		taskID:    taskID,
		doneChan:  doneChan,
		errChan:   errChan,
	}

	return doneChan, errChan
}
