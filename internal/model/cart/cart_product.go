package cart

//Link between cart and products, showing what products are in the cart
type CartProduct struct {
	id           int
	cartId       int
	productId    int
	productCount int
}

func NewCartProduct(cartId int, productId int, productCount int) CartProduct {
	return CartProduct{
		cartId:       cartId,
		productId:    productId,
		productCount: productCount,
	}
}

func (cp *CartProduct) SetProductCount(productCount int) {
	cp.productCount = productCount
}

func (cp *CartProduct) GetProductCount() int {
	return cp.productCount
}
