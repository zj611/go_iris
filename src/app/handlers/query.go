/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 下午4:06
 */
package handlers

import (
	"encoding/json"
	"example.com/m/src/app"
	"github.com/kataras/iris/v12"
)

func GetByName(a *app.App) func(ctx iris.Context) {
	return func(ctx iris.Context) {

		body, err := ctx.GetBody()
		if err != nil{
			panic("read body error!")
		}

		var tmp map[string]interface{}
		json.Unmarshal(body, &tmp)

		name := tmp["name"].(string)

		res,_ := a.QueryI.GetByName(name)

		WriteResult(ctx, map[string]interface{}{
			"items": res,
		})
	}
}

