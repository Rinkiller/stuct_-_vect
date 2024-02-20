package main

import (
	"fmt"
	"os"
	"time"
)

const (
	DateFotm = "2006-01-02 15:04:05"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
	delOldKey(t time.Time)
}

var _ Cache = (*cacheImpl)(nil)

func newCacheImpl(keys []string, val []string) *cacheImpl {
	if len(keys) != len(val) {
		os.Exit(125)
	}
	var cache map[string]string = make(map[string]string)
	var crd map[string]time.Time = make(map[string]time.Time)
	for i, k := range keys {
		cache[k] = val[i]
		crd[k] = time.Now()
	}
	return &cacheImpl{cache, crd}
}

type cacheImpl struct {
	cache map[string]string
	crd   map[string]time.Time
}

func (c *cacheImpl) Get(k string) (string, bool) {
	// TODO implement me
	if val, e := c.cache[k]; e {
		return val, true
	}
	return "", false
}

func (c *cacheImpl) Set(k, v string) {
	c.cache[k] = v
	c.crd[k] = time.Now()
}

func (c *cacheImpl) delOldKey(t time.Time) {
	for k, v := range c.crd {
		if t.After(v) {
			delete(c.cache, k)
			delete(c.crd, k)
		}
	}
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	if ok {
		return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
	}
	return "", false
}

type cacheImplLRU struct {
	cache  map[string]string
	crd    map[string]time.Time
	length int
}

var _ Cache = (*cacheImplLRU)(nil)

func (c *cacheImplLRU) delOldKey(t time.Time) {
	for k, v := range c.crd {
		if t.After(v) {
			delete(c.cache, k)
			delete(c.crd, k)
		}
	}
}
func newCacheImplLRU(keys []string, val []string, length int) *cacheImplLRU {
	if (len(keys) != len(val)) || (len(keys) > length) {
		os.Exit(125)
	}
	var cache map[string]string = make(map[string]string)
	var crd map[string]time.Time = make(map[string]time.Time)
	for i, k := range keys {
		cache[k] = val[i]
		crd[k] = time.Now()
	}
	return &cacheImplLRU{cache: cache, crd: crd, length: length}
}
func (c *cacheImplLRU) Get(k string) (string, bool) {
	// TODO implement me
	if val, e := c.cache[k]; e {
		return val, true
	}
	return "", false
}

func (c *cacheImplLRU) getOldK() string {
	var resultK string
	var ind int
	nowT := time.Now()
	for k, v := range c.crd {
		difference := nowT.Sub(v)
		if ind <= int(difference) {
			resultK = k
			ind = int(difference)
		}
	}
	return resultK
}

func (c *cacheImplLRU) Set(k, v string) {
	if _, ok := c.cache[k]; ok {
		if v == c.cache[k] {
			c.crd[k] = time.Now()
			return
		} else {
			c.cache[k] = v
			c.crd[k] = time.Now()
			return
		}

	} else if len(c.cache) < c.length {
		c.cache[k] = v
		c.crd[k] = time.Now()
		return
	} else {
		key := c.getOldK()
		delete(c.cache, key)
		delete(c.crd, key)
		c.cache[k] = v
		c.crd[k] = time.Now()
	}

}

func main() {
	a := []string{"one", "two", "three"}
	b := []string{"one day", "two days", "three days"}
	cs := newCacheImpl(a, b)
	db := newDbImpl(cs)
	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))
	fmt.Println(db.Get("two"))
	fmt.Println(db.Get("1"))
	fmt.Println(db.cache)
	time.Sleep(2 * time.Second)
	db.cache.Set("next", "next commit")
	fmt.Println(db.cache)
	db.cache.delOldKey(time.Now())
	time.Sleep(2 * time.Second)
	db.cache.Set("next", "next commit")
	fmt.Println(db.cache)

	c := newCacheImplLRU(a, b, 4)
	fmt.Println(c)
	fmt.Println(c.Get("one"))
	fmt.Println(c.Get("one1"))
	c.Set("2", "gzdfg")
	fmt.Println(c)
	time.Sleep(2 * time.Second)
	c.Set("2", "gzdfg")
	fmt.Println(c)
	c.Set("2", "gzdfg1111")
	fmt.Println(c)
	c.Set("4", "AAAAAAA")
	fmt.Println(c)
	c.Set("7", "777777")
	fmt.Println(c)
}
