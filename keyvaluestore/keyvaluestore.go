package keyvaluestore

import (
	"sync"
	"time"
)

// KeyValue represents a key-value pair with TTL.
type KeyValue struct {
	Value      string
	Expiration time.Time
}

// KeyValueStore represents the key-value store.
type KeyValueStore struct {
	data map[string]KeyValue
	mu   sync.RWMutex
}

// NewKeyValueStore creates a new instance of KeyValueStore.
func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: make(map[string]KeyValue),
	}
}

// Set adds or updates a key-value pair in the store with a specified TTL.
func (kv *KeyValueStore) Set(key, value string, ttl time.Duration) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	expiration := time.Now().Add(ttl)
	kv.data[key] = KeyValue{
		Value:      value,
		Expiration: expiration,
	}
}

// Get retrieves the value associated with a key from the store.
func (kv *KeyValueStore) Get(key string) (string, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	item, ok := kv.data[key]
	if !ok {
		return "", false
	}

	// Check if the item has expired
	if item.Expiration.IsZero() || time.Now().Before(item.Expiration) {
		return item.Value, true
	}

	// If the item has expired, remove it from the store
	delete(kv.data, key)
	return "", false
}
