package order

//Link between orders and products, contains products added to order
type OrderProduct struct {
	id           int
	orderId      int
	productId    int
	productCount int
}

func NewOrderProduct(orderId int, productId int, productCount int) OrderProduct {
	return OrderProduct{
		orderId:      orderId,
		productId:    productId,
		productCount: productCount,
	}
}

func (op *OrderProduct) GetOrderId() int {
	return op.orderId
}

func (op *OrderProduct) GetProductId() int {
	return op.productId
}

func (op *OrderProduct) SetProductCount(productCount int) {
	op.productCount = productCount
}

func (op *OrderProduct) GetProductCount() int {
	return op.productCount
}
