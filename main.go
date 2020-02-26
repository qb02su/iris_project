package main

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"io"
	"io/ioutil"
	"iriswork/config"
	"iriswork/controller"
	"iriswork/datasource"
	"iriswork/model"
	"iriswork/service"
	"iriswork/utils"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	app := newApp()

	//应用App设置
	configation(app)

	//路由设置
	mvcHandle(app)

	config := config.InitConfig()

	addr := ":" + config.Port
	app.Run(
		iris.Addr(addr),                               //在端口8080进行监听
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,                        //对json数据序列化更快的配置
	)
}

//构建App
func newApp() *iris.Application {
	app := iris.New()

	//设定应用图标
	app.Favicon("./static/favicons/favicon.ico")

	//设置日志级别  开发阶段为debug
	app.Logger().SetLevel("debug")

	//注册静态资源
	app.StaticWeb("/static", "./static")
	app.StaticWeb("/manage/static", "./static")
	app.StaticWeb("/img", "./uploads")

	//注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	})

	return app
}

/**
 * MVC 架构模式处理
 */
func mvcHandle(app *iris.Application) {

	//启用session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})

	//获取redis实例
	redis := datasource.NewRedis()
	//设置session的同步位置为redis
	sessManager.UseDatabase(redis)

	//实例化mysql数据库引擎
	engine := datasource.NewMysqlEngine()

	//管理员模块功能
	adminService := service.NewAdminService(engine)

	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessManager.Start,
	)
	admin.Handle(new(controller.AdminController))

	//用户功能模块
	userService := service.NewUserService(engine)
	user := mvc.New(app.Party("/v1/users"))
	user.Register(
		userService,
		sessManager.Start,
	)
	user.Handle(new(controller.UserController))
	//获取用户详细信息
	app.Get("/v1/user/{user_name}", func(context context.Context) {
		userName := context.Params().Get("user_name")
		var user model.User
		_, err := engine.Where(" user_name = ? ", userName).Get(&user)
		if err != nil {
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_USERINFO,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_USERINFO),
			})
		} else {
			context.JSON(user)
		}
	})
	//获取地址信息
	app.Get("/v1/addresse/{address_id}", func(context context.Context) {
		address_id := context.Params().Get("address_id")

		addressID, err := strconv.Atoi(address_id)
		if err != nil {
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERINFO,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERINFO),
			})
		}

		var address model.Address
		_, err = engine.Id(addressID).Get(&address)
		if err != nil {
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_ORDERINFO,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_ORDERINFO),
			})
		}
		//查询数据成功
		context.JSON(address)
	})

	//统计功能模块
	statisService := service.NewStatisService(engine)
	statis := mvc.New(app.Party("/statis/{model}/{date}/"))
	statis.Register(
		statisService,
		sessManager.Start,
	)
	statis.Handle(new(controller.StatisController))

	//订单模块
	orderService := service.NewOrderService(engine)
	order := mvc.New(app.Party("/bos/orders/"))
	order.Register(
		orderService,
		sessManager.Start,
	)
	order.Handle(new(controller.OrderController)) //控制器

	//商铺模块
	shopService := service.NewShopService(engine)
	shop := mvc.New(app.Party("/shopping/restaurants/"))
	shop.Register(
		shopService,
		sessManager.Start,
	)
	shop.Handle(new(controller.ShopController)) //控制器

	//添加食品类别
	categoryService := service.NewCategoryService(engine)
	category := mvc.New(app.Party("/shopping/"))
	category.Register(
		categoryService,
	)
	category.Handle(new(controller.CategoryController)) //控制器

	//食品模块
	foodService := service.NewFoodService(engine)
	foodMvc := mvc.New(app.Party("/shopping/v2/foods/"))
	foodMvc.Register(
		foodService,
	)
	foodMvc.Handle(new(controller.FoodController)) //控制器

	//文件上传
	app.Post("/admin/update/avatar/{adminId}", func(context context.Context) {
		adminId := context.Params().Get("adminId")
		iris.New().Logger().Info(adminId)

		file, info, err := context.FormFile("file")
		if err != nil {
			iris.New().Logger().Info(err.Error())
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}
		defer file.Close()
		fname := info.Filename

		out, err := os.OpenFile("./uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			iris.New().Logger().Info(err.Error())
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}
		iris.New().Logger().Info("文件路径：" + out.Name())
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}

		intAdminId, _ := strconv.Atoi(adminId)
		adminService.SaveAvatarImg(int64(intAdminId), fname)
		context.JSON(iris.Map{
			"status":     utils.RECODE_OK,
			"image_path": fname,
		})
	})

	//地址Poi检索
	app.Get("/v1/pois", func(context context.Context) {
		path := context.Request().URL.String()
		iris.New().Logger().Info(path)

		rs, err := http.Get("https://elm.cangdu.org" + path)
		if err != nil {
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_SEARCHADDRESS,
				"message": utils.Recode2Text(utils.RESPMSG_ERROR_SEARCHADDRESS),
			})
			return
		}

		//请求成功
		body, err := ioutil.ReadAll(rs.Body)
		var searchList []*model.PoiSearch
		//安马歇尔 马歇尔
		json.Unmarshal(body, &searchList)
		context.JSON(&searchList)
	})

	//上传图片
	app.Post("/v1/addimg/{model}", func(context context.Context) {
		model := context.Params().Get("model")
		iris.New().Logger().Info(model)

		file, info, err := context.FormFile("file")
		if err != nil {
			iris.New().Logger().Info(err.Error())
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}
		defer file.Close()
		fname := info.Filename

		//判断上传的目录是否存在，如果不存在的话，先创建目录
		isExist, err := utils.PathExists("./uploads/" + model)

		if !isExist {
			err := os.Mkdir("./uploads/"+model, 0777)
			if err != nil {
				context.JSON(iris.Map{
					"status":  utils.RECODE_FAIL,
					"type":    utils.RESPMSG_ERROR_PICTUREADD,
					"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
				})
				return
			}
		}

		out, err := os.OpenFile("./uploads/"+model+"/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			iris.New().Logger().Info(err.Error())
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}
		iris.New().Logger().Info("文件路径：" + out.Name())
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			context.JSON(iris.Map{
				"status":  utils.RECODE_FAIL,
				"type":    utils.RESPMSG_ERROR_PICTUREADD,
				"failure": utils.Recode2Text(utils.RESPMSG_ERROR_PICTUREADD),
			})
			return
		}

		//上传成功
		context.JSON(iris.Map{
			"status":     utils.RECODE_OK,
			"image_path": fname,
		})
	})
}

/**
 * 项目设置
 */
func configation(app *iris.Application) {

	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    " not found ",
			"data":   iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    " interal error ",
			"data":   iris.Map{},
		})
	})
}
