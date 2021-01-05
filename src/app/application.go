/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 上午10:05
 */
package app

import (
	"example.com/m/pkg/mysql"
	"example.com/m/pkg/storage"
	"example.com/m/src/app/config"
	"example.com/m/src/app/services"
	"fmt"
	"github.com/kataras/iris/v12"
)

type App struct {
	*iris.Application
	DbCheck pinger
	Config  *config.Config
	QueryI  services.QueryI
	InputI  services.InputI
}

type Configurator func(*App)

type pinger func() bool

func (app *App) Configure(configurators ...Configurator) *App {
	for _, cfg := range configurators {
		cfg(app)
	}
	return app
}

func NewApp() *App {
	app := iris.Default()
	cfg := config.ReadEnv()

	fmt.Printf("s%", cfg.MysqlURL)

	db := mysql.NewGormDB(&cfg.MysqlURL)
	db.SingularTable(true)

	app.Logger().Info("begin auto migration")
	if err := db.AutoMigrate(
		&storage.TestTable{},
	).Error; err != nil {
		panic(err)
	}

	app.Logger().Info("auto migration success")

	inputI := services.NewInputI(db)
	queryI := services.NewQueryI(db)
	return &App{
		Application: app,
		DbCheck: func() bool {
			if err := db.DB().Ping(); err != nil {
				app.Logger().Error("db ping error:", err)
				return false
			}
			return true
		},
		Config: cfg,
		QueryI: queryI,
		InputI: inputI,
	}
}
