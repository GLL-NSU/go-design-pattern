package structural

import "fmt"

type Phone interface {
	Call()
}
type Decorator struct {
	phone Phone
}

func (d *Decorator) Call() {
	d.phone.Call()
}

type IPhone struct {
	Decorator
}

func (i *IPhone) Call() {
	fmt.Println("使用iPhone拨打电话")
}

type IPhoneDecorator struct {
	Decorator
}

func (i *IPhoneDecorator) Call() {
	i.phone.Call()
	fmt.Println("使用iPhoneDecorator装饰器后可以进行录音")
}

func NewIPhoneDecorator(phone Phone) Phone {
	return &IPhoneDecorator{Decorator{phone}}
}

func UseDecorator() {
	fmt.Println("--------开始：装饰模式--------")
	fmt.Println("非装饰器模式进行通话")
	iphone := &IPhone{}
	iphone.Call()
	fmt.Println("使用装饰器模式进行通话")
	NewIPhoneDecorator(iphone).Call()
	fmt.Println("--------结束：装饰模式--------")
	fmt.Println()
}
