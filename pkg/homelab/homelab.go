package homelab

import (
	"fmt"
)

type Service interface {
	Identifier() string
	Task(taskID string) (Task, bool)
}

type Task interface {
	Identifier() string
}

type Bridge interface {
	Listen(ctx Context) error
}

type HomeLab struct {
	bridge   Bridge
	services map[string]Service
}

func New(bridge Bridge, services []Service) HomeLab {
	lab := HomeLab{
		bridge:   bridge,
		services: make(map[string]Service),
	}

	for _, service := range services {
		serviceID := service.Identifier()
		lab.services[serviceID] = service
	}

	return lab
}

func (hl HomeLab) Listen() error {
	ctx := NewContext()
	go func() {
		for qt := range ctx.taskQueue {
			service, ok := hl.services[qt.serviceID]
			if !ok {
				qt.errChan <- fmt.Errorf("unknown service: %s", qt.serviceID)
				continue
			}

			task, ok := service.Task(qt.taskID)
			if !ok {
				qt.errChan <- fmt.Errorf("unknown task: %s", qt.taskID)
				continue
			}

			fmt.Printf("Executing %s:%s\n", service.Identifier(), task.Identifier())
			qt.doneChan <- true
		}
	}()
	return hl.bridge.Listen(ctx)
}
