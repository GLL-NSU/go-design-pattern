package singleton

import "fmt"

type lazySingleton struct {
}

var lazyInstance *lazySingleton

func GetLazyInstance() *lazySingleton {
	if lazyInstance == nil {
		lazyInstance = new(lazySingleton)
	}
	return lazyInstance
}
func (s *lazySingleton) DoSomething() {
	fmt.Println("懒汉式单例对象正在执行业务逻辑")
}
