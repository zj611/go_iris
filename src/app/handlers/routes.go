/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 上午10:09
 */
package handlers

import (
	"example.com/m/src/app"
	"fmt"
)

func RegisterRoutes(app *app.App) {
	fmt.Println("RegisterRoutes...")

	input := app.Party("/input")
	{
		input.Post("/report", Report(app))

		input.Post("/update", Update(app))
	}

	query := app.Party("/query")
	{
		query.Post("/getByName", GetByName(app))

		//query.Post("/update", Update(app))
	}
}
