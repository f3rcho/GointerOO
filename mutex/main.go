package main

import (
	"fmt"
	"sync"
)

var (
	data   map[string]string
	rwLock sync.RWMutex
)

func readData(key string) string {
	rwLock.RLock()
	defer rwLock.RUnlock()
	return data[key]
}

func writeData(key, value string) {
	rwLock.Lock()
	defer rwLock.Unlock()
	data[key] = value
}

func main() {
	data = make(map[string]string)

	// Writing data concurrently
	for i := 0; i < 10; i++ {
		go writeData(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	// Reading data concurrently
	for i := 0; i < 10; i++ {
		go func(index int) {
			key := fmt.Sprintf("key%d", index)
			value := readData(key)
			fmt.Println("Read:", key, value)
		}(i)
	}
}
