package maps

import "sync"

/*
目前 Go 还没有正式发布泛型特性，我们还不能实现一个通用的支持泛型的加锁 map。但是，将要发布的泛型方案已经可以验证测试了，离发布也不远了，
也许发布之后 sync.Map 就支持泛型了。当然了，如果没有泛型支持，我们也能解决这个问题。我们可以通过 interface{}来模拟泛型，
但还是要涉及接口和具体类型的转换，比较复杂，还不如将要发布的泛型方案更直接、性能更好。

缺点：
虽然使用读写锁可以提供线程安全的 map，但是在大量并发读写的情况下，锁的竞争会非
常激烈。锁是性能下降的万恶之源之一。
*/
type RWLockMap struct {
	m    map[interface{}]interface{}
	lock sync.RWMutex
}

func CreateRWLockMap() *RWLockMap {
	m := make(map[interface{}]interface{}, 0)
	return &RWLockMap{m: m}
}

func (m *RWLockMap) Get(key interface{}) (interface{}, bool) {
	m.lock.RLock()
	v, ok := m.m[key]
	m.lock.RUnlock()
	return v, ok
}

func (m *RWLockMap) Set(key interface{}, value interface{}) {
	m.lock.Lock()
	m.m[key] = value
	m.lock.Unlock()
}

func (m *RWLockMap) Del(key interface{}) {
	m.lock.Lock()
	delete(m.m, key)
	m.lock.Unlock()
}

/*
---------------------------------------------------------------------------------------
示例2
*/
type RWMap struct { // 一个读写锁保护的线程安全的map
	sync.RWMutex // 读写锁保护下面的map字段
	m            map[int]int
}

// 新建一个RWMap
func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}

func (m *RWMap) Get(k int) (int, bool) { //从map中读取一个值
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k] // 在锁的保护下从map中读取
	return v, existed
}

func (m *RWMap) Set(k int, v int) { // 设置一个键值对
	m.Lock() // 锁保护
	defer m.Unlock()
	m.m[k] = v
}

func (m *RWMap) Delete(k int) { //删除一个键
	m.Lock() // 锁保护
	defer m.Unlock()
	delete(m.m, k)
}

func (m *RWMap) Len() int { // map的长度
	m.RLock() // 锁保护
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap) Each(f func(k, v int) bool) { // 遍历map
	m.RLock() //遍历期间一直持有读锁
	defer m.RUnlock()
	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
