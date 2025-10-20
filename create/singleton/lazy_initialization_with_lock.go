package singleton

import (
	"fmt"
	"sync"
)

var lock = sync.Mutex{}

type lazyLockSingleton struct {
}

var lazyLockInstance *lazyLockSingleton

func GetLazyLockInstance() *lazyLockSingleton {
	lock.Lock()
	defer lock.Unlock()
	if lazyLockInstance == nil {
		lazyLockInstance = new(lazyLockSingleton)
	}
	return lazyLockInstance
}
func (s *lazyLockSingleton) DoSomething() {
	fmt.Println("懒汉式线程安全单例对象正在执行业务逻辑")
}
