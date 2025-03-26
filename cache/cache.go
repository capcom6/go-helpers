package cache

import (
	"sync"
	"time"
)

type item[T any] struct {
	value      T
	validUntil time.Time
}

func newItem[T any](value T, opts options) *item[T] {
	item := &item[T]{
		value:      value,
		validUntil: opts.validUntil,
	}

	return item
}

func (i *item[T]) isExpired(now time.Time) bool {
	return !i.validUntil.IsZero() && now.After(i.validUntil)
}

type Cache[T any] struct {
	items map[string]*item[T]
	ttl   time.Duration
	empty T

	mux sync.RWMutex
}

// New creates a new Cache instance with the specified configuration.
//
// T is the type of the values stored in the cache.
//
// Returns a pointer to a new Cache instance.
func New[T any](cfg Config) *Cache[T] {
	return &Cache[T]{
		items: make(map[string]*item[T]),
		ttl:   cfg.TTL,
	}
}

func (c *Cache[T]) newItem(value T, opts ...Option) *item[T] {
	o := options{}
	if c.ttl > 0 {
		o.validUntil = time.Now().Add(c.ttl)
	}
	o.apply(opts...)

	return newItem(value, o)
}

// Set sets the value for the given key in the cache.
func (c *Cache[T]) Set(key string, value T, opts ...Option) error {
	c.mux.Lock()
	c.items[key] = c.newItem(value, opts...)
	c.mux.Unlock()

	return nil
}

// SetOrFail is like Set, but returns ErrKeyExists if the key already exists.
func (c *Cache[T]) SetOrFail(key string, value T, opts ...Option) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	if _, ok := c.items[key]; ok {
		return ErrKeyExists
	}

	c.items[key] = c.newItem(value, opts...)

	return nil
}

func (c *Cache[T]) getItem(getter func() (*item[T], bool)) (*item[T], error) {
	item, ok := getter()

	if !ok {
		return nil, ErrKeyNotFound
	}

	if item.isExpired(time.Now()) {
		return nil, ErrKeyExpired
	}

	return item, nil
}

func (c *Cache[T]) getValue(getter func() (*item[T], bool)) (T, error) {
	item, err := c.getItem(getter)
	if err != nil {
		return c.empty, err
	}

	return item.value, nil
}

// Get gets the value for the given key from the cache.
//
// If the key is not found, it returns ErrKeyNotFound.
// If the key has expired, it returns ErrKeyExpired.
// Otherwise, it returns the value and nil.
func (c *Cache[T]) Get(key string) (T, error) {
	return c.getValue(func() (*item[T], bool) {
		c.mux.RLock()
		item, ok := c.items[key]
		c.mux.RUnlock()

		return item, ok
	})
}

// GetAndDelete is like Get, but also deletes the key from the cache.
func (c *Cache[T]) GetAndDelete(key string) (T, error) {
	return c.getValue(func() (*item[T], bool) {
		c.mux.Lock()
		item, ok := c.items[key]
		delete(c.items, key)
		c.mux.Unlock()

		return item, ok
	})
}

// Delete removes the item associated with the given key from the cache.
// If the key does not exist, it performs no action and returns nil.
// The operation is safe for concurrent use.
func (c *Cache[T]) Delete(key string) error {
	c.mux.Lock()
	delete(c.items, key)
	c.mux.Unlock()

	return nil
}

// Drain returns a map of all the non-expired items in the cache.
// The returned map is a snapshot of the cache at the time of the call.
// The cache is cleared after the call.
// The operation is safe for concurrent use.
func (c *Cache[T]) Drain() map[string]T {
	t := time.Now()

	c.mux.Lock()
	copy := c.items
	c.items = make(map[string]*item[T], len(copy))
	c.mux.Unlock()

	items := make(map[string]T, len(copy))
	for key, item := range copy {
		if item.isExpired(t) {
			continue
		}
		items[key] = item.value
	}

	return items
}

// Cleanup removes all expired items from the cache.
// The operation is safe for concurrent use.
func (c *Cache[T]) Cleanup() {
	t := time.Now()

	c.mux.Lock()
	defer c.mux.Unlock()

	for key, item := range c.items {
		if item.isExpired(t) {
			delete(c.items, key)
		}
	}
}
