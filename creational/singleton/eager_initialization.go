package singleton

import "fmt"

type eagerSingleton struct {
}

var eagerInstance = new(eagerSingleton)

func GetEagerInstance() *eagerSingleton {
	return eagerInstance
}
func (s *eagerSingleton) DoSomething() {
	fmt.Println("饿汉式单例对象正在执行业务逻辑")
}
