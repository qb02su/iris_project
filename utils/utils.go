package utils

import (
	"fmt"
	"github.com/google/logger"
	"github.com/kataras/iris"
	"os"
	"reflect"
	"time"
)

//根据Json格式设置obj对象
func SetObjByJson(obj interface{}, data map[string]interface{}) error {
	for key, value := range data {
		if err := setField(obj, key, value); err != nil {
			logger.Error("SetObjByJson set field fail.")
			return err
		}
	}
	return nil
}

//设置结构体中的变量
func setField(obj interface{}, name string, value interface{}) error {
	structData := reflect.TypeOf(obj).Elem()
	fieldValue, result := structData.FieldByName(name)
	if !result {
		logger.Error("No such field ", name)
		return fmt.Errorf("No such field %s", name)
	}

	//结构体中变量的类型
	fieldType := fieldValue.Type
	//参数的值
	val := reflect.ValueOf(value)
	//参数的类型
	valTypeStr := val.Type().String()
	//结构体中变量的类型
	fieldTypeStr := fieldType.String()
	//float64 to int
	if valTypeStr == "float64" && fieldTypeStr == "int" {
		val = val.Convert(fieldType)
	}

	//类型必须匹配
	if fieldType != val.Type() {
		return fmt.Errorf("value type %s didn't match obj field type %s ", valTypeStr, fieldTypeStr)
	}

	//fieldValue.Set(val)

	return nil
}

func LogInfo(app *iris.Application, v ...interface{}) {
	app.Logger().Info(v)
}

func LogError(app *iris.Application, v ...interface{}) {
	app.Logger().Error(v)
}

func LogDebug(app *iris.Application, v ...interface{}) {
	app.Logger().Debug(v)
}

/**
 * 格式化数据
 */
func FormatDatetime(time time.Time) string {
	return time.Format("2006-01-02 03:04:05")
}

/**
 * 判断某个路径是否存在
 * 返回两个值：第一个值为路径是否存在；第二个值返回error
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
