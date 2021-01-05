/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/15 下午5:38
 */
package main

import (
	"example.com/m/src/app"
	"example.com/m/src/app/handlers"
	"fmt"
	"github.com/kataras/iris/v12"
)

func main()  {

	fmt.Println("hello")

	//app := iris.New()
	//app.Get("/", GetPageList(app))(iris.Addr(":8081"))
	//app.Run
	server := app.NewApp()

	server.Configure(handlers.RegisterRoutes)

	server.Run(iris.Addr(":8081"))
}



