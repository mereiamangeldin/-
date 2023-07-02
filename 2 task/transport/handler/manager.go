package handler

import "github.com/mereiamangeldin/OnePlusTask2/service"

type Manager struct {
	srv *service.Manager
}

func NewManager(srv *service.Manager) *Manager {
	return &Manager{srv: srv}
}
