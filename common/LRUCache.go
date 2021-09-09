package common

import (
	"container/list"
	"urlShortner/constant"
)

var KeyCacheMap map[string]*LRUCache

type LRUCache struct {
	cap int                      // capacity
	l   *list.List               // doubly linked list
	m   map[string]*list.Element // hash table for checking if list node exists
}

type CacheData struct {
	ShortUrl  string
	ActualUrl string
}

func init() {
	KeyCacheMap = make(map[string]*LRUCache)
}

func InitializeCache() *LRUCache {
	return &LRUCache{
		cap: constant.CACHE_CAPACITY,
		l:   new(list.List),
		m:   make(map[string]*list.Element, constant.CACHE_CAPACITY),
	}
}

func GetCache(key string) *LRUCache {
	if cache, present := KeyCacheMap[key]; present {
		return cache
	}
	newCache := InitializeCache()
	KeyCacheMap[key] = newCache

	return newCache
}

func (c *LRUCache) Get(shortURL string) (string, bool) {
	// check if list node exists
	if node, ok := c.m[shortURL]; ok {
		val := node.Value.(*list.Element).Value.(CacheData).ActualUrl
		// move node to front
		c.l.MoveToFront(node)
		return val, true
	}
	return "", false
}

func (c *LRUCache) Put(shortURL string, actualURL string) {
	// check if list node exists
	if node, ok := c.m[shortURL]; ok {
		// move the node to front
		c.l.MoveToFront(node)
		// update the value of a list node
		node.Value.(*list.Element).Value = CacheData{ShortUrl: shortURL, ActualUrl: actualURL}
	} else {
		// delete the last list node if the list is full
		if c.l.Len() == c.cap {
			// get the key that we want to delete
			idx := c.l.Back().Value.(*list.Element).Value.(CacheData).ShortUrl
			// delete the node pointer in the hash map by key
			delete(c.m, idx)
			// remove the last list node
			c.l.Remove(c.l.Back())
		}
		// initialize a list node
		node := &list.Element{
			Value: CacheData{
				ShortUrl:  shortURL,
				ActualUrl: actualURL,
			},
		}
		// push the new list node into the list
		ptr := c.l.PushFront(node)
		// save the node pointer in the hash map
		c.m[shortURL] = ptr
	}
}
