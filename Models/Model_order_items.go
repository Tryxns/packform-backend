package Models

type Orders_items struct {
	Id             int `json:"id"`
	Order_id       int `json:"order_id"`
	Price_per_unit int `json:"price_per_unit"`
	Quantity       int `json:"quantity"`
	Product        int `json:"product"`
}

func (b *Orders_items) TableName() string {
	return "Orders_items"
}
