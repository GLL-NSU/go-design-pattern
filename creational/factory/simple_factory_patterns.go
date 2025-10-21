package factory

import "fmt"

// Fruit1 Fruit 水果类的抽象接口
type Fruit1 interface {
	PrintFruitName()
}

// Apple1 Apple 苹果具体实现类，实现了水果接口
type Apple1 struct {
	Name string
}

func (a *Apple1) PrintFruitName() {
	fmt.Println("通过简单工厂生产的apple")
}

// Orange1 Orange 橘子具体实现类，实现了水果接口
type Orange1 struct {
	Name string
}

func (o *Orange1) PrintFruitName() {
	fmt.Println("通过简单工厂生产的orange")
}

// SimpleFruitFactory 水果简单工厂类
type SimpleFruitFactory struct {
}

func (f *SimpleFruitFactory) CreateFruit(name string) Fruit1 {
	if name == "apple" {
		return &Apple1{Name: name}
	} else if name == "orange" {
		return &Orange1{Name: name}
	}
	return nil
}

// UseSimpleFactoryPattern 使用简单工厂模式
func UseSimpleFactoryPattern() {
	fmt.Println("--------开始：简单工厂模式--------")
	simpleFactory := &SimpleFruitFactory{}
	apple := simpleFactory.CreateFruit("apple")
	apple.PrintFruitName()
	orange := simpleFactory.CreateFruit("orange")
	orange.PrintFruitName()
	fmt.Println("--------结束：简单工厂模式--------")
	fmt.Println()
}
