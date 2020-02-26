package model

/**
 * 订单状态结构体定义
 */
type OrderStatus struct {
	StatusId   int64  `xorm:"pk autoincr" json:"id"` //主键
	StatusDesc string `xorm:"varchar(255)"`          // 订单状态描述
}
