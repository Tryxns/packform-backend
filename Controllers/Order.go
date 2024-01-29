package Controllers

import (
	"net/http"
	"packform/Models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type jsondata struct {
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
	Count  int64                `json:"count"`
	Data   []Models.ResultOrder `json:"data"`
}

func GetOrders(c *gin.Context) {
	var orders []Models.ResultOrder
	qry_limit := c.Query("limit")
	qry_offset := c.Query("offset")
	// filter_product := c.Query("product")
	// filter_order := c.Query("order")
	search := c.Query("search")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")

	var limit int
	if qry_limit == "" {
		limit = 5
	} else {
		limit, _ = strconv.Atoi(qry_limit)
	}

	var offset int
	if qry_offset == "" {
		offset = 0
	} else {
		offset, _ = strconv.Atoi(qry_offset)
	}

	count, err := Models.GetOrders(&orders, limit, offset, search, start_date, end_date)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, jsondata{
			limit, offset, count, orders,
		})
	}
}
