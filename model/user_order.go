package model

import "time"

/**
 * 用户订单结构实体定义
 */
type UserOrder struct {
	Id            int64        `xorm:"pk autoincr" json:"id"` //主键
	SumMoney      int64        `xorm:"default 0" json:"sum_money"`
	Time          time.Time    `xorm:"DateTime" json:"time"`         //时间
	OrderTime     uint64       `json:"order_time"`                   //订单创建时间
	OrderStatusId int64        `xorm:"index" json:"order_status_id"` //订台状态id
	OrderStatus   *OrderStatus `xorm:"-"` //订单对象

	UserId        int64        `xorm:"index" json:"user_id"`         //用户编号Id
	User          *User        `xorm:"-"`                            //订单对应的账户，并不进行结构体字段映射

	ShopId        int64        `xorm:"index" json:"shop_id"`         //用户购买的商品编号
	Shop          *Shop        `xorm:"-"`                            //商品结构体，不进行映射

	AddressId     int64        `xorm:"index" json:"address_id"`      //地址结构体的Id
	Address       *Address     `xorm:"-"`                            //地址结构体，不进行映射
	DelFlag       int64        `xorm:"default 0" json:"del_flag"`    //删除标志 0为正常 1为已删除
}

/**
 * 查询得到的userOrder实体转换为resp的json格式
 */
func (this *UserOrder) UserOrder2Resp() interface{} {
	respDesc := map[string]interface{}{
		"id":                   this.Id,
		"total_amount":         this.SumMoney,
		"user_id":              this.User.UserName,          //用户名
		"status":               this.OrderStatus.StatusDesc, //订单状态
		"restaurant_id":        this.Shop.ShopId,            //商铺id
		"restaurant_image_url": this.Shop.ImagePath,         //商铺图片
		"restaurant_name":      this.Shop.Name,              //商铺名称
		"formatted_created_at": this.Time,
		"status_code":          0,
		"address_id":           this.Address.AddressId, //订单地址id
	}

	statusDesc := map[string]interface{}{
		"color":     "f60",
		"sub_title": "15分钟内支付",
		"title":     this.OrderStatus.StatusDesc,
	}
	
	respDesc["status_bar"] = statusDesc
	return respDesc
}
