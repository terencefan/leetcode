package p901

type IndexPrice struct {
	index int
	price int
}

type StockSpanner struct {
	index  int
	prices []IndexPrice
}

func Constructor() StockSpanner {
	return StockSpanner{0, []IndexPrice{{0, 1 << 31}}}
}

func (this *StockSpanner) Next(price int) int {
	this.index++

	for this.prices[len(this.prices)-1].price <= price {
		this.prices = this.prices[:len(this.prices)-1]
	}
	defer func() { this.prices = append(this.prices, IndexPrice{this.index, price}) }()

	return this.index - this.prices[len(this.prices)-1].index
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
