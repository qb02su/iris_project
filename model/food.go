package model

/**
 * 食品结构体定义
 */
type Food struct {
	Id           int64         `xorm:"pk autoincr" json:"item_id"` //食品id
	Name         string        `json:"name"`                       //食品名称
	Description  string        `json:"description"`                //食品描述
	Rating       int           `json:"rating"`                     //食品评分
	MonthSales   int           `json:"month_sales"`                //月销量
	ImagePath    string        `json:"image_path"`                 //食品图片地址
	Activity     string        `json:"activity"`                   //食品活动
	Attributes   string        `json:"attributes"`                 //食品特点
	Specs        string        `json:"specs"`                      //食品规格
	CategoryId   int64         `xorm:"index"`                      //食品种类ID
	Category     *FoodCategory `xorm:"-"`                          //食品种类
	RestaurantId int64         `xorm:"index"`                      //商铺ID
	Restaurant   *Shop         `xorm:"-"`                          //食品店铺信息
	DelFlag      int           `json:"del_flag"`                   //是否已经被删除 0表示未删除 1表示1被删除
}
