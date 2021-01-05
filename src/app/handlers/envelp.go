/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 下午2:11
 */
package handlers
import "github.com/kataras/iris/v12"

func WriteResult(ctx iris.Context, r interface{}) {
	ctx.JSON(map[string]interface{}{"result": r})
}