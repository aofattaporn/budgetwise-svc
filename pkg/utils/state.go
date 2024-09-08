package utils

import "sync"

type state struct {
	isError bool
	mu      *sync.RWMutex
}

type IState interface {
	IsErr() bool
	SetErr()
}

func InitState() IState {
	return &state{
		isError: false,
		mu:      &sync.RWMutex{},
	}
}

func (s *state) IsErr() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isError
}

func (s *state) SetErr() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.isError = true
}

type batchState struct {
	isRunning bool
	mu        *sync.Mutex
}

type IBatchState interface {
	Start() (isSuccess bool)
	End()
}

func InitBatchState() IBatchState {
	return &batchState{
		isRunning: false,
		mu:        &sync.Mutex{},
	}
}

func (s *batchState) Start() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.isRunning {
		return false
	}
	s.isRunning = true
	return true
}

func (s *batchState) End() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.isRunning = false
}
