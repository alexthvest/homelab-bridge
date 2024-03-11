package stablediffusion

import (
	"github.com/alexthvest/homelab-bridge/pkg/homelab"
)

type Service struct {
	tasks map[string]homelab.Task
}

func NewService(tasks ...homelab.Task) *Service {
	service := Service{
		tasks: make(map[string]homelab.Task),
	}

	for _, task := range tasks {
		taskID := task.Identifier()
		service.tasks[taskID] = task
	}

	return &service
}

func (s *Service) Identifier() string {
	return "stable-diffusion"
}

func (s *Service) Task(taskID string) (homelab.Task, bool) {
	task, ok := s.tasks[taskID]
	return task, ok
}
