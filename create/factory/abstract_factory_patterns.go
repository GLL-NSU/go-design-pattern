package factory

import "fmt"

// AbstractApple 抽象苹果接口
type AbstractApple interface {
	PrintAppleName()
}

// AbstractOrange 抽象橘子接口
type AbstractOrange interface {
	PrintOrangeName()
}

// AbstractFactory 抽象工厂接口
type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateOrange() AbstractOrange
}

// CNApple 中国苹果
type CNApple struct {
}

func (a *CNApple) PrintAppleName() {
	fmt.Println("通过抽象工厂模式创建了中国苹果")
}

// CNOrange 中国橘子
type CNOrange struct {
}

func (o *CNOrange) PrintOrangeName() {
	fmt.Println("通过抽象工厂模式创建了中国橘子")
}

// CNFactory 中国工厂
type CNFactory struct {
}

func (f *CNFactory) CreateApple() AbstractApple {
	return &CNApple{}
}
func (f *CNFactory) CreateOrange() AbstractOrange {
	return &CNOrange{}
}

// USApple 美国苹果
type USApple struct {
}

func (a *USApple) PrintAppleName() {
	fmt.Println("通过抽象工厂模式创建美国苹果")
}

// USOrange 美国橘子
type USOrange struct {
}

func (o *USOrange) PrintOrangeName() {
	fmt.Println("通过抽象工厂模式创建美国橘子")
}

// USFactory 美国工厂
type USFactory struct {
}

func (f *USFactory) CreateApple() AbstractApple {
	return &USApple{}
}
func (f *USFactory) CreateOrange() AbstractOrange {
	return &USOrange{}
}
func UseAbstractFactoryPattern() {
	fmt.Println("--------开始：抽象工厂模式--------")
	cnFactory := &CNFactory{}
	cnApple := cnFactory.CreateApple()
	cnApple.PrintAppleName()
	cnOrange := cnFactory.CreateOrange()
	cnOrange.PrintOrangeName()
	usFactory := &USFactory{}
	usApple := usFactory.CreateApple()
	usApple.PrintAppleName()
	usOrange := usFactory.CreateOrange()
	usOrange.PrintOrangeName()
	fmt.Println("--------结束：抽象工厂模式--------")
	fmt.Println()
}
