//go:build wireinject

package main

import (
	"e-marketing/internal/ioc"
	"e-marketing/internal/job"
	"e-marketing/internal/repository"
	"e-marketing/internal/repository/dao"
	"e-marketing/internal/service"

	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		// third-party
		ioc.InitSendgrid,
		ioc.InitDB,

		// dao 部分
		dao.NewCursorDAO,
		dao.NewRecipientDAO,
		dao.NewShopDAO,
		dao.NewSenderDAO,
		dao.NewScenarioDAO,

		// repository 部分
		repository.NewCursorRepo,
		repository.NewRecipientRepo,
		repository.NewSenderRepo,
		repository.NewScenarioRepo,
		repository.NewShopRepo,

		// service 部分
		service.NewShopService,
		service.NewEmailService,
		service.NewEmailSendService,

		// job 部分
		ioc.InitJobs,
		job.NewUninstalledJob,

		// server
		ioc.InitWebServer,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
