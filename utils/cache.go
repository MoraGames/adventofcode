package utils

func NewCache[T comparable]() map[T]int {
	return make(map[T]int)
}

func CacheAdd[T comparable](cache map[T]int, key T, amount int) int {
	if _, ok := cache[key]; !ok {
		cache[key] = 0
	}
	cache[key] += amount
	return cache[key]
}

func CacheSub[T comparable](cache map[T]int, key T, amount int, stopAtZero bool) int {
	if currAmount, ok := cache[key]; ok {
		if stopAtZero && currAmount < amount {
			cache[key] = 0
		} else {
			cache[key] -= currAmount
		}
	}
	return cache[key]
}

func CacheDelete[T comparable](cache map[T]int, key T) {
	delete(cache, key)
}

func CacheExists[T comparable](cache map[T]int, key T) bool {
	_, ok := cache[key]
	return ok
}
