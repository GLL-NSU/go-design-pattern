package factory

import "fmt"

// Fruit2 水果类的抽象接口
type Fruit2 interface {
	PrintFruitName()
}

// Apple2  苹果具体实现类，实现了水果接口
type Apple2 struct {
	Name string
}

func (a *Apple2) PrintFruitName() {
	fmt.Println("通过工厂方法生产的apple")
}

// Orange2  橘子具体实现类，实现了水果接口
type Orange2 struct {
	Name string
}

func (o *Orange2) PrintFruitName() {
	fmt.Println("通过工厂方法生产的orange")
}

// AbstractFruitFactory 水果工厂抽象类
type AbstractFruitFactory interface {
	CreateFruit(name string) Fruit2
}

// AppleFactory 苹果工厂类
type AppleFactory struct {
}

func (f *AppleFactory) CreateFruit(name string) Fruit2 {
	return &Apple2{Name: name}
}

// OrangeFactory 橘子工厂类
type OrangeFactory struct {
}

func (f *OrangeFactory) CreateFruit(name string) Fruit2 {
	return &Orange2{Name: name}
}

// UseFactoryMethodPattern 工厂方法模式使用示例
func UseFactoryMethodPattern() {
	fmt.Println("--------开始：工厂方法模式--------")
	appleFactory := &AppleFactory{}
	apple := appleFactory.CreateFruit("apple")
	apple.PrintFruitName()
	orangeFactory := &OrangeFactory{}
	orange := orangeFactory.CreateFruit("orange")
	orange.PrintFruitName()
	fmt.Println("--------结束：工厂方法模式--------")
	fmt.Println()
}
