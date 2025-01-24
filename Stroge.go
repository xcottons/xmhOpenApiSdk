package xmhOpenApiSdk

import "sync"

type Storage interface {
	Save(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

type DefaultMemoryStorage struct {
	mStorage sync.Map
}

func (d *DefaultMemoryStorage) Save(key string, value interface{}) error {
	d.mStorage.Store(key, value)
	return nil
}

func (d *DefaultMemoryStorage) Get(key string) (interface{}, error) {
	if value, ok := d.mStorage.Load(key); ok {
		return value, nil
	}
	return nil, nil
}
