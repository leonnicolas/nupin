package store

var _ Store = &MemoryStore{}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[string]struct{})}
}

type MemoryStore struct {
	data map[string]struct{}
}

func (m *MemoryStore) Exist(key string) (bool, error) {
	if _, ok := m.data[key]; ok {
		return true, nil
	}
	return false, nil
}

func (m *MemoryStore) Set(key string) error {
	m.data[key] = struct{}{}
	return nil
}

func (m *MemoryStore) Delete(key string) error {
	delete(m.data, key)
	return nil
}
