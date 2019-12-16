package store

// MemoryStore is the memory store
// NOTE: really shouldn't be used in prod
type MemoryStore map[string]interface{}

// Destroy removes a session from the store
func (store *MemoryStore) Destroy(sid string) error {
	delete(*store, sid)
	return nil
}

// Get retrieves the session from the store
func (store *MemoryStore) Get(sid string) (interface{}, error) {
	return (*store)[sid], nil
}

// Set sets a session with a given value
func (store *MemoryStore) Set(sid string, v interface{}) error {
	(*store)[sid] = v
	return nil
}
