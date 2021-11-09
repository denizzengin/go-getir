package main

import (
	"fmt"
)

var customMap *InMemoryMap

// Set : In-memory rest api set endpoint
func Set(key string, val interface{}) {
	customMap.Mutex.Lock()
	customMap.KeyValuePair[key] = fmt.Sprint(val)
	customMap.Mutex.Unlock()
}

// Get : In-memory rest api get endpoint
func Get(key string) (string, error) {
	customMap.Mutex.Lock()
	defer customMap.Mutex.Unlock()
	if val, ok := customMap.KeyValuePair[key]; ok {
		return val, nil
	}
	return "", ErrNotFound
}
