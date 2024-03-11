package main

import (
	"github.com/alexthvest/homelab-bridge/pkg/homelab"
	"github.com/alexthvest/homelab-bridge/pkg/services/stablediffusion"
	"github.com/alexthvest/homelab-bridge/pkg/services/stablediffusion/txt2img"
)

func setupServices() []homelab.Service {
	services := make([]homelab.Service, 0)

	sdService := stablediffusion.NewService(
		txt2img.Task{},
	)
	services = append(services, sdService)

	return services
}
