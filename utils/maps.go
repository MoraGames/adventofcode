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

func NewCheck[T comparable]() map[T]struct{} {
	return make(map[T]struct{})
}

func CheckInsert[T comparable](check map[T]struct{}, value T) {
	check[value] = struct{}{}
}

func CheckTryInsert[T comparable](check map[T]struct{}, value T) bool {
	if _, ok := check[value]; !ok {
		check[value] = struct{}{}
		return true
	}
	return false
}

func CheckDelete[T comparable](check map[T]struct{}, value T) {
	delete(check, value)
}

func CheckExists[T comparable](check map[T]struct{}, value T) bool {
	_, ok := check[value]
	return ok
}
