// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"gin-admin/internal/app/api"
	"gin-admin/internal/app/dao"
	"gin-admin/internal/app/module/adapter"
	"gin-admin/internal/app/router"
	"gin-admin/internal/app/service"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		dao.DaoSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		service.ServiceSet,
		api.APISet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
