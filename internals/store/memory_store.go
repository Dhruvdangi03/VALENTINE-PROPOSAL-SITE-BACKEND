package store

import (
	"Valentine-Proposal-Site-Backend/internals/models"
	"sync"
)

type MemoryStore struct {
	data map[string]models.ProposalDTO
	mu   sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]models.ProposalDTO),
	}
}

func (m *MemoryStore) Save(short string, proposal models.ProposalDTO) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[short] = proposal
}

func (m *MemoryStore) Get(short string) (models.ProposalDTO, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.data[short]
	return val, ok
}
