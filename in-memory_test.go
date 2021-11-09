package main

import (
	"sync"
	"testing"
)

func TestInMemorySet(t *testing.T) {
	customMap = &InMemoryMap{KeyValuePair: make(map[string]string), Mutex: &sync.Mutex{}}
	Set("dummyKey", "dummyValue")
	if k, ok := customMap.KeyValuePair["dummyKey"]; !ok {
		t.Fail()
	} else if k != "dummyValue" {
		t.Fail()
	}
}

func TestInMemoryGet(t *testing.T) {
	customMap = &InMemoryMap{KeyValuePair: make(map[string]string), Mutex: &sync.Mutex{}}
	Set("dummyKeyGet", "dummyValue")
	v, e := Get("dummyKeyGet")
	if e != nil {
		t.Fail()
	}

	if v != "dummyValue" {
		t.Fail()
	}
}
