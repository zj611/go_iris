/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 上午10:17
 */
package handlers

import (
	"encoding/json"
	"example.com/m/pkg/storage"
	"example.com/m/src/app"
	"github.com/kataras/iris/v12"
	"strconv"
	"time"
)

func Report(a *app.App) func(ctx iris.Context) {
	return func(ctx iris.Context) {

		body, err := ctx.GetBody()
		if err != nil{
			panic("read body error!")
		}
		InsertData := new(storage.TestTable)
		var tmp map[string]interface{}
		json.Unmarshal([]byte(body), &tmp)
		InsertData.Name = tmp["name"].(string)
		t := tmp["age"].(string)
		InsertData.Age,_ = strconv.Atoi(t)
		InsertData.Sex = tmp["sex"].(string)
		InsertData.MobilePhone = tmp["mobile_phone"].(string)

		math := tmp["math"].(string)
		math_, _ := strconv.ParseFloat(math, 32)
		InsertData.Math = float32(math_)

		english := tmp["english"].(string)
		english_, _ := strconv.ParseFloat(english, 32)
		InsertData.English = float32(english_)

		chinese := tmp["chinese"].(string)
		chinese_, _ := strconv.ParseFloat(chinese, 32)
		InsertData.Chinese = float32(chinese_)

		InsertData.BirthDay,_ = time.ParseInLocation("2006-01-02",  tmp["birthday"].(string), time.Local)


		res := a.InputI.Insert(InsertData)
		//if err != nil {
		//	WriteResult(ctx, map[string]interface{}{
		//		"items": "GetPageList()处理错误！",
		//	})
		//	return
		//}
		//fmt.Println(result)

		WriteResult(ctx, map[string]interface{}{
			"items": res,
		})
	}
}



func Update(a *app.App) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		body, err := ctx.GetBody()
		if err != nil{
			panic("read body error!")
		}



		var tmp map[string]interface{}
		json.Unmarshal(body, &tmp)

		updateData := new(storage.ReqUpdateData)

		updateData.Name = tmp["name"].(string)

		math := tmp["math"].(string)
		math_, _ := strconv.ParseFloat(math, 32)
		updateData.Math = float32(math_)

		english := tmp["english"].(string)
		english_, _ := strconv.ParseFloat(english, 32)
		updateData.English = float32(english_)

		chinese := tmp["chinese"].(string)
		chinese_, _ := strconv.ParseFloat(chinese, 32)
		updateData.Chinese = float32(chinese_)

		testTable := &storage.TestTable{
			Name: tmp["name"].(string),
			Math: float32(math_),
			Chinese: float32(chinese_),
			English: float32(english_),
		}


		res := a.InputI.Update(testTable)

		WriteResult(ctx, map[string]interface{}{
			"items": res,
		})
	}
}


func GetItemArray(parent map[string]interface{}, key string) []interface{} {
	val, ok := parent[key].([]interface{})
	if ok {
		return val
	} else {
		return nil
	}
}
func GetItemMap(parent map[string]interface{}, key string) map[string]interface{} {
	val, ok := parent[key].(map[string]interface{})
	if ok {
		return val
	} else {
		return nil
	}
}
func GetItemString(parent map[string]interface{}, key string) string {
	val, ok := parent[key].(string)
	if ok {
		return val
	}
	return ""
}
func GetItemFloat64(parent map[string]interface{}, key string) float64 {
	val, ok := parent[key].(float64)
	if ok {
		return val
	}
	return 0
}