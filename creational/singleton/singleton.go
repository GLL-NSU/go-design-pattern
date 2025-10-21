package singleton

import "fmt"

func UseSingletonPattern() {
	fmt.Println("--------开始：单例模式--------")
	GetEagerInstance().DoSomething()
	GetLazyInstance().DoSomething()
	GetLazyLockInstance().DoSomething()
	GetLazyAtomicSingleton().DoSomething()
	GetLazyOnceSingleton().DoSomething()
	fmt.Println("--------结束：单例模式--------")
	fmt.Println()
}
