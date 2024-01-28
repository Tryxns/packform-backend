package Models

type Orders struct {
	Id          int    `json:"id"`
	Created_at  string `json:"created_at"`
	Customer_id string `json:"customer_id"`
	Order_name  string `json:"order_name"`
}

type ResultOrder struct {
	Id                 int     `json:"id"`
	Order_name         string  `json:"order_name"`
	Company_name       string  `json:"company_name"`
	Customer_name      string  `json:"customer_name"`
	Created_at         string  `json:"created_at"`
	Delivered_quantity int     `json:"delivered_quantity"`
	Quantity           int     `json:"quantity"`
	Price_per_unit     float32 `json:"price_per_unit"`
}

func (b *Orders) TableName() string {
	return "packform.orders"
}
