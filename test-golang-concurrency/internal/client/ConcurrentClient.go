package client

import (
	"hash/fnv"
	"sync"
)

// Number of shards available to us. Mocking a "sharded" database logic.
const numShards = 32

type DataEntry struct {
	data string
	mu   sync.RWMutex
}

type Shard struct {
	mu   sync.RWMutex
	data map[string]*DataEntry
}

type ConcurrentClient struct {
	shards []*Shard
}

type IConcurrentClient interface {
	ReadData(key string) string
	WriteData(key string, data string) error
}

func NewConcurrentClient() IConcurrentClient {
	shards := make([]*Shard, numShards)
	for i := range shards {
		shards[i] = &Shard{data: make(map[string]*DataEntry)}
	}

	return &ConcurrentClient{
		shards: shards,
	}
}

func (c *ConcurrentClient) findShard(key string) *Shard {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return c.shards[hash.Sum32()%numShards]
}

func (c *ConcurrentClient) ReadData(key string) string {
	shard := c.findShard(key)

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	entry, ok := shard.data[key]
	if !ok {
		return ""
	} else {
		entry.mu.RLock()
		defer entry.mu.RUnlock()

		return entry.data
	}
}

func (c *ConcurrentClient) WriteData(key string, data string) error {
	shard := c.findShard(key)
	entry := shard.getOrCreateEntry(key)

	entry.mu.Lock()
	defer entry.mu.Unlock()
	entry.data = data

	return nil
}

func (s *Shard) getOrCreateEntry(key string) *DataEntry {
	// Quickly check if key already exists
	s.mu.RLock()
	entry, ok := s.data[key]
	s.mu.RUnlock()

	// If key doesn't exist, add new entry
	if !ok {
		// Lock to avoid duplicate updates
		s.mu.Lock()
		defer s.mu.Unlock()

		// Set outside of if statement to avoid weird variable scoping
		entry, ok = s.data[key]

		// While locked, check if another system updated the key since the last lock
		if !ok {
			entry = &DataEntry{
				mu: sync.RWMutex{},
			}
			s.data[key] = entry
		}
	}

	// Return entry, whether retrieved or created
	return entry
}
