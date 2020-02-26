package service

import (
	"iriswork/model"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

/**
 * 管理员服务
 * 标准的开发模式将每个实体的提供的功能以接口标准的形式定义,供控制层进行调用。
 *
 */
type AdminService interface {
	//通过管理员用户名+密码 获取管理员实体 如果查询到，返回管理员实体，并返回true
	//否则 返回 nil ，false
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)
	GetByAdminId(adminId int64) (model.Admin, bool)
	//获取管理员总数
	GetAdminCount() (int64, error)
	SaveAvatarImg(adminId int64, fileName string) bool
	GetAdminList(offset, limit int) []*model.Admin
}

func NewAdminService(db *xorm.Engine) AdminService {
	return &adminSevice{
		engine: db,
	}
}

/**
 * 管理员的服务实现结构体
 */
type adminSevice struct {
	engine *xorm.Engine
}

/**
 * 查询管理员总数
 */
func (ac *adminSevice) GetAdminCount() (int64, error) {
	count, err := ac.engine.Count(new(model.Admin))

	if err != nil {
		panic(err.Error())
		return 0, err
	}

	return count, nil
}

/**
 * 通过用户名和密码查询管理员
 */
func (ac *adminSevice) GetByAdminNameAndPassword(username, password string) (model.Admin, bool) {
	var admin model.Admin

	ac.engine.Where(" admin_name = ? and pwd = ? ", username, password).Get(&admin)

	return admin, admin.AdminId != 0
}

/**
 * 查询管理员信息
 */
func (ac *adminSevice) GetByAdminId(adminId int64) (model.Admin, bool) {
	var admin model.Admin

	ac.engine.Id(adminId).Get(&admin)

	return admin, admin.AdminId != 0
}

/**
 * 保存头像信息
 */
func (ac *adminSevice) SaveAvatarImg(adminId int64, fileName string) bool {
	admin := model.Admin{Avatar: fileName}
	_, err := ac.engine.Id(adminId).Cols(" avatar ").Update(&admin)
	return err != nil
}

/**
 * 获取管理员列表
 * offset：获取管理员的便宜量
 * limit：请求管理员的条数
 */
func (ac *adminSevice) GetAdminList(offset, limit int) []*model.Admin {
	var adminList []*model.Admin

	err := ac.engine.Limit(limit, offset).Find(&adminList)
	if err != nil {
		iris.New().Logger().Error(err.Error())
		panic(err.Error())
		return nil
	}
	return adminList
}
