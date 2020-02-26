package datasource

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"iriswork/config"
)

/**
 * 返回Redis实例
 */
func NewRedis() *redis.Database {

	var database *redis.Database

	//项目配置
	cmsConfig := config.InitConfig()
	if cmsConfig != nil {
		iris.New().Logger().Info(" hello ")
		rd := cmsConfig.Redis
		iris.New().Logger().Info(rd)
		database = redis.New(service.Config{
			Network:     rd.NetWork,
			Addr:        rd.Addr + ":" + rd.Port,
			Password:    rd.Password,
			Database:    "",
			MaxIdle:     0,
			MaxActive:   10,
			IdleTimeout: service.DefaultRedisIdleTimeout,
			Prefix:      rd.Prefix,
		})
	} else {
		iris.New().Logger().Info(" hello  error ")
	}
	return database
}
