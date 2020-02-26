package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iriswork/service"
	"iriswork/utils"
	"strconv"
)

/**
 * 食品模块控制器
 */
type FoodController struct {
	Ctx     iris.Context
	Service service.FoodService
}

/**
 * url：foods/count?
 * type：Get
 * desc：获取所有的食品记录总数
 */
func (fc *FoodController) GetCount() mvc.Result {
	iris.New().Logger().Info(" 食品记录总数 ")
	result, err := fc.Service.GetFoodCount()

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RESPMSG_FAIL,
				"count":  0,
			},
		}
	}

	//查询成功
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RESPMSG_OK,
			"count":  result,
		},
	}
}

/**
 * url：/shopping/v2/foods/
 * type：type
 * desc：请求查询食品列表数据
 */
func (fc *FoodController) Get() mvc.Result {
	offset, err := strconv.Atoi(fc.Ctx.Params().Get("offset"))
	limit, err := strconv.Atoi(fc.Ctx.Params().Get("limit"))
	if err != nil {
		offset = 0
		limit = 20
	}
	list, err := fc.Service.GetFoodList(offset, limit)
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_FOODLIST,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_FOODLIST),
			},
		}
	}
	//成功
	return mvc.Response{
		Object: &list,
	}
}
