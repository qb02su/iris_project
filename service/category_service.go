package service

import (
	"github.com/go-xorm/xorm"
	"iriswork/model"
	"github.com/kataras/iris"
)

/**
 * 食品种类服务接口
 */
type CategoryService interface {
	AddCategory(model *model.FoodCategory) bool
	GetCategoryByShopId(shopId int64) ([]model.FoodCategory, error)
	GetAllCategory() ([]model.FoodCategory, error)

	//获取商铺信息操作
	GetRestaurantInfo(shop_id int64) (model.Shop, error)

	//保存记录操作
	SaveFood(food model.Food) bool
	SaveShop(shop model.Shop) bool

	//删除记录操作
	DeleteShop(restaurantId int) bool
	DeleteFood(foodId int) bool
}

/**
 * 种类服务实现结构体
 */
type categoryService struct {
	Engine *xorm.Engine
}

/**
 * 实例化种类服务:服务器
 */
func NewCategoryService(engine *xorm.Engine) CategoryService {
	return &categoryService{
		Engine: engine,
	}
}

/**
 * 获取某个商铺的信息
 */
func (cs *categoryService) GetRestaurantInfo(shopId int64) (model.Shop, error) {
	var shop model.Shop
	_, err := cs.Engine.Id(shopId).Get(&shop)
	return shop, err
}

/**
 * 通过商铺Id获取食品类型
 */
func (cs *categoryService) GetCategoryByShopId(shopId int64) ([]model.FoodCategory, error) {
	categories := make([]model.FoodCategory, 0)
	err := cs.Engine.Where(" restaurant_id = ? ", shopId).Find(&categories)
	return categories, err
}

/**
 * 添加食品种类记录
 */
func (cs *categoryService) AddCategory(category *model.FoodCategory) bool {
	_, err := cs.Engine.Insert(category)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 * 获取所有的种类信息
 */
func (cs *categoryService) GetAllCategory() ([]model.FoodCategory, error) {
	categories := make([]model.FoodCategory, 0)
	err := cs.Engine.Where(" parent_category_id = ? ", 0).Find(&categories)
	return categories, err
}

/**
 * 保存食品记录
 */
func (cs *categoryService) SaveFood(food model.Food) bool {
	_, err := cs.Engine.Insert(&food)
	return err == nil
}

/**
 * 保存商户记录
 */
func (cs *categoryService) SaveShop(shop model.Shop) bool {

	_, err := cs.Engine.Insert(&shop)

	if err != nil {
		iris.New().Logger().Info(err.Error())
	}

	return err == nil
}

/**
 * 删除商铺
 */
func (cs *categoryService) DeleteShop(restaurantId int) bool {
	shop := model.Shop{ShopId: restaurantId, Dele: 1}

	_, err := cs.Engine.Where(" shop_id = ? ", restaurantId).Cols("dele").Update(&shop)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}
	return err == nil
}

/**
 * 删除食品
 */
func (cs *categoryService) DeleteFood(foodId int) bool {

	food := model.Food{Id: int64(foodId), DelFlag: 1}

	_, err := cs.Engine.Where(" id = ? ", foodId).Cols("del_flag").Update(&food)
	if err != nil {
		iris.New().Logger().Info(err.Error())
	}

	return err == nil
}
