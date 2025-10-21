package main

import (
	"go-design-pattern/creational/factory"
	"go-design-pattern/creational/singleton"
	"go-design-pattern/structural"
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
	// 代理模式
	structural.UseProxy()

}
