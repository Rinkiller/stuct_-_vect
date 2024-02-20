package main

import (
	"fmt"
	"os"
	"time"
)

type cacheImplLRU struct {
	cache  map[string]string
	crd    map[string]time.Time
	length int
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
		fmt.Println("DEBUG!!!!!")
		fmt.Println(key)
		delete(c.cache, key)
		delete(c.crd, key)
		c.cache[k] = v
		c.crd[k] = time.Now()
	}

}

func main() {
	a := []string{"one", "two", "three"}
	b := []string{"one day", "two days", "three days"}
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
