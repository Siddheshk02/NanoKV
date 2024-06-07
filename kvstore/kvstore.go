package kvstore

import "sync"

type KeyValueStore struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: make(map[string]string),
	}
}

func (kv *KeyValueStore) Set(key, value string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[key] = value
}

func (kv *KeyValueStore) Get(key string) (string, bool) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	val, ok := kv.data[key]

	return val, ok
}

func (kv *KeyValueStore) Delete(key string) bool {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	_, ok := kv.data[key]
	if !ok {
		return ok
	}

	delete(kv.data, key)

	return ok
}
