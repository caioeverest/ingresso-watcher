package repository

import (
	"sync"

	"github.com/caioeverest/ingresso-watcher/service/errors"
)

type Memory struct {
	mu    sync.Mutex
	items map[string]string
}

func NewMemory() *Memory {
	return &Memory{items: make(map[string]string)}
}

func (m *Memory) Set(id string, name string) {
	m.mu.Lock()
	m.items[id] = name
	m.mu.Unlock()
}

func (m *Memory) GetById(id string) (name string, found bool) {
	m.mu.Lock()
	name, found = m.items[id]
	m.mu.Unlock()
	return
}

func (m *Memory) GetAll() (items map[string]string) {
	m.mu.Lock()
	items = m.items
	m.mu.Unlock()
	return
}

func (m *Memory) Delete(id string) error {
	m.mu.Lock()
	if _, ok := m.items[id]; !ok {
		return errors.NotFound
	}
	delete(m.items, id)
	m.mu.Unlock()
	return nil
}
