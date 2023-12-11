package memo

import "sync"

// 并发、不重复、无阻塞的 cache
// 尽管 entry 中的 res 会被多个 goroutine 访问，当并不需要互斥锁；
// ready channel 的关闭一定会发生在其他 goroutine 接收到广播事件之前，因为第一个 goroutine 对这些变量的写操作是一定发生在其他的读操作之前，不会发生数据竞争

// Func is the type of the function to memoize.
type Func func(key string) (any, error)

type result struct {
	value any
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // Closed when res is ready
}

// Memo caches the results of calling Func.
type Memo struct {
	f   Func
	mtx sync.Mutex // guards cache
	// cache map[string]result
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value any, err error) {
	memo.mtx.Lock()
	e := memo.cache[key]
	if e == nil {
		// 首次请求：缓存计算结果并广播 ready 条件
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mtx.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // broadcast ready condition
	} else {
		// 重复请求
		memo.mtx.Unlock()
		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}
