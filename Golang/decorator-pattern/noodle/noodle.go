package noodle

type Noodle interface {
	getPrice() int
}

type VegetableMania struct {
}

func (v *VegetableMania) getPrice() int {
	return 15
}

type Bonus interface {
	Noodle
	BuyTwoBonusOne() int
}

type CheeseTopping struct {
	noodle Noodle
	price  int
}

func (c *CheeseTopping) getPrice() int {
	c.price = c.noodle.getPrice() + 7
	return c.price
}

func (c *CheeseTopping) BuyTwoBonusOne() int {
	return c.price * 2
}

type TomatoTopping struct {
	noodle Noodle
	price  int
}

func (t *TomatoTopping) getPrice() int {
	t.price = t.noodle.getPrice() + 10
	return t.price
}

func (t *TomatoTopping) BuyTwoBonusOne() int {
	return t.price * 2
}
