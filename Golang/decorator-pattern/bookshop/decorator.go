package bookshop

type Book interface {
	GetPrice() float64
}

type BookImpl struct {
	Title string
	Price float64
}

func (b *BookImpl) GetPrice() float64 {
	return b.Price
}

type DiscountedBook interface {
	Book
	SetDiscount(discount float64)
}

type DiscountedBookImpl struct {
	Book     Book
	Discount float64
}

func (d *DiscountedBookImpl) GetPrice() float64 {
	return d.Book.GetPrice() * (1.0 - d.Discount)
}

func (d *DiscountedBookImpl) SetDiscount(discount float64) {
	d.Discount = discount
}
