package model

/**
 * 用户订单详情结构体
 */
type OrderDetail struct {
	UserOrder   `xorm:"extends"`
	User        `xorm:"extends"`
	OrderStatus `xorm:"extends"`
	Shop        `xorm:"extends"`
	Address     `xorm:"extends"`
}


func (detail *OrderDetail) OrderDetail2Resp() interface{} {
	respDesc := map[string]interface{}{
		"id":                   detail.UserOrder.Id,
		"total_amount":         detail.UserOrder.SumMoney,
		"user_id":              detail.User.UserName,          //用户名
		"status":               detail.OrderStatus.StatusDesc, //订单状态
		"restaurant_id":        detail.Shop.ShopId,            //商铺id
		"restaurant_image_url": detail.Shop.ImagePath,         //商铺图片
		"restaurant_name":      detail.Shop.Name,              //商铺名称
		"formatted_created_at": detail.Time,
		"status_code":          0,
		"address_id":           detail.Address.AddressId, //订单地址id
	}

	statusDesc := map[string]interface{}{
		"color":     "f60",
		"sub_title": "15分钟内支付",
		"title":     detail.OrderStatus.StatusDesc,
	}
	respDesc["status_bar"] = statusDesc
	return respDesc
}
