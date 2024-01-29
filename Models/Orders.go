package Models

import (
	"fmt"
	"packform/Config"
)

func GetAllOrders(orders *[]Orders) (err error) {
	if err = Config.DB.Debug().Find(orders).Error; err != nil {
		return err
	}
	return nil
}

func GetOrders(results *[]ResultOrder, limit int, offset int, search string, start_date string, end_date string) (count int64, err error) {
	count = 0
	query_string := Config.DB.Table("packform.orders").Debug().
		Select(
			"orders.id, order_name, customers.name as \"customer_name\", company_name, orders.created_at, delivered_quantity, quantity, price_per_unit").
		Joins(
			"inner join packform.customers on orders.customer_id = customers.user_id").
		Joins(
			"inner join packform.customer_companies on customers.company_id = customer_companies.company_id").
		Joins(
			"inner join packform.order_items on orders.id = order_items.order_id").
		Joins(
			"inner join packform.deliveries on order_items.id = deliveries.order_item_id")

	if start_date != "" && end_date != "" {
		query_string.Where("orders.created_at between ? and ?", start_date, end_date)
	}
	// if order != "" {
	// 	query_string.Where("LOWER(order_name) like ?", fmt.Sprintf("%%%s%%", order))
	// }
	// if product != "" {
	// 	query_string.Where("LOWER(order_items.product) like ?", fmt.Sprintf("%%%s%%", product))
	// }
	if search != "" {
		query_string.Where(
			"LOWER(order_name) like ? or LOWER(order_items.product) like ?",
			fmt.Sprintf("%%%s%%", search), fmt.Sprintf("%%%s%%", search),
		)
	}
	query_string.Count(&count)
	query_string.Limit(limit).Offset(offset).Scan(results)

	return count, nil
}
