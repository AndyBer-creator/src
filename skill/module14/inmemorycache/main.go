package main

import (
	"sync"
	"time"
)

// CacheEntry структура для хранения значения и времени установки
type CacheEntry struct {
	SettledAt time.Time
	Value     interface{}
}

// Cache интерфейс для работы с кешем
type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

// InMemoryCache структура для кеша в памяти
type InMemoryCache struct {
	sync.RWMutex
	cache        map[string]*CacheEntry
	expiration   time.Duration
	cleanupTimer *time.Timer
}

// NewInMemoryCache конструктор для создания нового экземпляра InMemoryCache
func NewInMemoryCache(expireIn time.Duration) *InMemoryCache {
	cache := &InMemoryCache{
		cache:      make(map[string]*CacheEntry),
		expiration: expireIn,
	}
	cache.startCleanupTimer()
	return cache
}

// startCleanupTimer запускает таймер для очистки устаревших записей
func (c *InMemoryCache) startCleanupTimer() {
	c.cleanupTimer = time.AfterFunc(c.expiration, c.cleanup)
}

// cleanup очищает все записи, срок действия которых истек
func (c *InMemoryCache) cleanup() {
	c.Lock()
	defer c.Unlock()

	now := time.Now()
	for k, v := range c.cache {
		if now.Sub(v.SettledAt) > c.expiration {
			delete(c.cache, k)
		}
	}

	// Перезапускаем таймер для следующей чистки
	c.startCleanupTimer()
}

// Set устанавливает новое значение в кэш
func (c *InMemoryCache) Set(key string, value interface{}) {
	c.Lock()
	defer c.Unlock()

	entry := &CacheEntry{
		SettledAt: time.Now(),
		Value:     value,
	}
	c.cache[key] = entry
}

// Get получает значение из кеша
func (c *InMemoryCache) Get(key string) interface{} {
	c.RLock()
	defer c.RUnlock()

	if entry, found := c.cache[key]; found && time.Since(entry.SettledAt) <= c.expiration {
		return entry.Value
	}
	return nil
}

func main() {
	cache := NewInMemoryCache(time.Second * 10)

	cache.Set("key1", "value1")

	time.Sleep(time.Second * 5)

	value := cache.Get("key1") // Должен вернуть "value1"
	println(value.(string))

	time.Sleep(time.Second * 7)

	value = cache.Get("key1") // Должен вернуть nil, потому что запись устарела
	println(value)
}
