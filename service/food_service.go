package service

import (
	"github.com/go-xorm/xorm"
	"iriswork/model"
)

type FoodService interface {
	//返回食品记录的总记录数
	GetFoodCount() (int64, error)
	//返回所有的食品列表数据
	GetFoodList(offset, limit int) ([]model.Food, error)
}

//food服务实体的实现结构体
type foodService struct {
	Engine *xorm.Engine
}

/**
 * 构建FoodService服务实例
 */
func NewFoodService(engine *xorm.Engine) FoodService {
	return &foodService{
		Engine: engine,
	}
}

/**
 * 获取食品总记录数
 */
func (fs *foodService) GetFoodCount() (int64, error) {
	count, err := fs.Engine.Where(" del_flag = 0 ").Count(new(model.Food))
	return count, err
}

/**
 * 获取食品列表数据
 */
func (fs *foodService) GetFoodList(offset, limit int) ([]model.Food, error) {
	foodList := make([]model.Food, 0)
	err := fs.Engine.Where(" del_flag  = 0 ").Limit(limit, offset).Find(&foodList)
	return foodList, err
}
