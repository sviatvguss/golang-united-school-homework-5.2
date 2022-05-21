package cache

import "time"

type Cache struct {
	values    map[string]string
	deadlines map[string]time.Time
}

func NewCache() Cache {
	return Cache{
		values:    map[string]string{},
		deadlines: map[string]time.Time{}}
}

func (c *Cache) Get(key string) (string, bool) {
	v, vOk := c.values[key]
	if !vOk {
		return "", false
	}
	d, dOk := c.deadlines[key]
	if dOk {
		if d.Before(time.Now()) {
			delete(c.values, key)
			delete(c.deadlines, key)
			return "", false
		}
	}
	return v, vOk
}

func (c *Cache) Put(key, value string) {
	c.values[key] = value
	_, ok := c.deadlines[key]
	if ok {
		delete(c.deadlines, key)
	}
}

func (c *Cache) Keys() []string {
	var keys = make([]string, 0, len(c.values))
	for key := range c.values {
		keys = append(keys, key)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.values[key] = value
	c.deadlines[key] = deadline
}
