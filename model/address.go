package model

/**
 * 地址结构实体
 */
type Address struct {
	AddressId     int64  `xorm:"pk autoincr" json:"id"`
	Address       string `json:"address"`        //地址
	Phone         string `json:"phone"`          //联系人手机号
	AddressDetail string `json:"address_detail"` //地址详情
	IsValid       int    `json:"is_valid"`
}
