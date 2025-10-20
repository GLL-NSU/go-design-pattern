package main

import (
	"go-design-pattern/create/factory"
	"go-design-pattern/create/singleton"
)

func main() {
	// 简单工厂模式
	factory.UseSimpleFactoryPattern()
	// 工厂方法模式
	factory.UseFactoryMethodPattern()
	// 抽象工厂模式
	factory.UseAbstractFactoryPattern()
	// 单例模式
	singleton.UseSingletonPattern()

}
