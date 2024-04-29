package main

// Path: cache_test.go
// generate test cache for Get Set function

import (
	"encoding/json"
	"testing"
)

type Data struct {
	Name string
	Age  int
}

func (d Data) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

func TestRedisCache_SetString(t *testing.T) {
	// generate test cache for Get Set function
	cache := NewRedisCache()
	// define key, value pari n a test table
	tests := []struct {
		key   string
		value string
	}{
		{"key", "value"},
		{"key1", "value1"},
		{"key2", "value2"},
	}
	for _, tt := range tests {
		err := cache.Set(tt.key, tt.value, 0)
		if err != nil {
			t.Errorf("failed to set key-value pair: %v", err)
		}
		value, err := cache.Get(tt.key)
		if err != nil {
			t.Errorf("failed to get value: %v", err)
		}
		if value != tt.value {
			t.Errorf("expected value is %v, got %v", tt.value, value)
		}
	}

	// test case when the value is a struct or other data type than native data type

	data := Data{"name", 20}
	err := cache.Set("data", data, 0)
	if err != nil {
		t.Errorf("failed to set key-value pair: %v", err)
	}
	value, err := cache.Get("data")
	if err != nil {
		t.Errorf("failed to get value: %v", err)
	}
	// convert value to Data struct then compare
	var d Data

	err = json.Unmarshal([]byte(value.(string)), &d)
	if err != nil {
		t.Errorf("failed to unmarshal value: %v", err)
	}
	if d != data {
		t.Errorf("expected value is %v, got %v", data, d)
	}
}
