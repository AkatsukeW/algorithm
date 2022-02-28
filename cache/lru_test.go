package cache

import (
	"math/rand"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(5)
	cache.Set(10, 20)
	cache.Set(20, 20)
	cache.Set(30, 20)
	cache.Set(40, 20)
	cache.Set(50, 20)
	cache.Set(60, 20)
	val, err := cache.Get(10)
	if err != nil {
		t.Errorf("get val error is: %s", err)
	}
	if val != nil {
		t.Errorf("cache miss err is: %s", err)
	}
	v, _ := cache.Get(60)
	println(v.(int))
}

func BenchmarkLRU(b *testing.B) {
	cache := NewLRUCache(8092)

	trace := make([]int64, b.N*2)
	for i := 0; i < b.N*2; i++ {
		trace[i] = rand.Int63() * 32768
	}
	b.ResetTimer()

	var hit, miss int
	for i := 0; i < b.N*2; i++ {
		if i%2 == 0 {
			_ = cache.Set(trace[i], trace[i])
		} else {
			_, err := cache.Get(trace[i])
			if err != nil {
				miss++
			} else {
				hit++
			}
		}
	}
	b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}
