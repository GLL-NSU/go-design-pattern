package singleton

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var lc = sync.Mutex{}
var initialized uint32

type lazyAtomicSingleton struct {
}

var lazyAtomicInstance *lazyAtomicSingleton

func GetLazyAtomicSingleton() *lazyAtomicSingleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return lazyAtomicInstance
	}
	lc.Lock()
	defer lc.Unlock()
	if initialized == 0 {
		lazyAtomicInstance = new(lazyAtomicSingleton)
		atomic.StoreUint32(&initialized, 1)
	}
	return lazyAtomicInstance
}
func (s *lazyAtomicSingleton) DoSomething() {
	fmt.Println("懒汉式高性能线程安全单例对象正在执行业务逻辑")
}
