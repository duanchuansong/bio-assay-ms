// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"bio-assay-ms/internal/app/api"
	"bio-assay-ms/internal/app/dao"
	"bio-assay-ms/internal/app/module/adapter"
	"bio-assay-ms/internal/app/router"
	"bio-assay-ms/internal/app/service"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		dao.RepoSet,
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
