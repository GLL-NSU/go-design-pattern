package structural

import "fmt"

type Goods struct {
	Kind string
	Fact bool
}

// Shopping 购物抽象层
type Shopping interface {
	Buy(goods *Goods)
}

type ChinaShopping struct {
}

func (c *ChinaShopping) Buy(goods *Goods) {
	fmt.Println("在中国进行购物", goods)
}

type ProxyShopping struct {
	Shopping Shopping
}

func (p *ProxyShopping) Buy(goods *Goods) {
	fact := checkFact(goods)
	if !fact {
		fmt.Println("商品真品检查失败", goods)
		return
	}
	p.Shopping.Buy(goods)
	checkCustoms(goods)
}

func checkFact(goods *Goods) bool {
	return goods.Fact
}
func checkCustoms(goods *Goods) {
	fmt.Println("商品海关检查通过", goods)
}

func NewProxy(shopping Shopping) Shopping {
	return &ProxyShopping{
		Shopping: shopping,
	}
}

func UseProxy() {
	fmt.Println("--------开始：代理模式--------")
	fmt.Println("使用非代理模式购买")
	chinaShopping := &ChinaShopping{}
	goods := &Goods{Kind: "电脑", Fact: true}
	if checkFact(goods) {
		chinaShopping.Buy(goods)
		checkCustoms(goods)
	}
	fmt.Println("使用代理模式购买")
	NewProxy(&ChinaShopping{}).Buy(goods)
	fmt.Println("--------结束：代理模式--------")
	fmt.Println()
}
