package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type lazyOnceSingleton struct {
}

var lazyOnceInstance *lazyOnceSingleton

func GetLazyOnceSingleton() *lazyOnceSingleton {
	once.Do(func() {
		lazyOnceInstance = new(lazyOnceSingleton)
	})
	return lazyOnceInstance
}
func (s *lazyOnceSingleton) DoSomething() {
	fmt.Println("懒汉式高性能简约线程安全单例对象正在执行业务逻辑")
}
