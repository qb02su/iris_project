package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"iriswork/service"
	"strings"
	"github.com/kataras/iris/mvc"
	"iriswork/utils"
)

/**
 * 统计功能控制者
 */
type StatisController struct {
	//上下文环境对象
	Ctx iris.Context

	//统计功能的服务实现接口
	Service service.StatisService

	//session
	Session *sessions.Session
}

var (
	ADMINMODULE = "ADMIN_"
	USERMODULE  = "USER_"
	ORDERMODULE = "ORDER_"
)

/**
 * 解析统计功能路由请求
 */
func (sc *StatisController) GetCount() mvc.Result {

	// /statis/user/2019-03-10/count

	path := sc.Ctx.Path()

	var pathSlice []string
	if path != "" {
		//  /statis/user/2019-03-10/count
		//  "" "statis "user" "2019-03-10" "count"
		pathSlice = strings.Split(path, "/")
	}

	//不符合请求格式
	if len(pathSlice) != 5 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	//将最前面的去掉
	// "statis "user" "2019-03-10" "count"
	pathSlice = pathSlice[1:]
	model := pathSlice[1]
	date := pathSlice[2]
	var result int64
	switch model {
	case "user":
		userResult := sc.Session.Get(USERMODULE + date)
		if userResult != nil {
			userResult = userResult.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  userResult,
				},
			}
		} else {
			iris.New().Logger().Error(date) //时间
			result = sc.Service.GetUserDailyCount(date)
			//设置缓存
			// date： 2019-04-23
			//  USER_2019-04-23 result
			sc.Session.Set(USERMODULE+date, result)
		}
	case "order":

		orderStatis := sc.Session.Get(ORDERMODULE + date)

		if orderStatis != nil {
			orderStatis = orderStatis.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  orderStatis,
				},
			}
		} else {
			result = sc.Service.GetOrderDailyCount(date)
			sc.Session.Set(ORDERMODULE+date, result)
		}

	case "admin":
		adminStatis := sc.Session.Get(ADMINMODULE + date)
		if adminStatis != nil {
			adminStatis = adminStatis.(float64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  adminStatis,
				},
			}
		} else {
			result = sc.Service.GetAdminDailyCount(date)
			sc.Session.Set(ADMINMODULE, result)
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  result,
		},
	}
}
