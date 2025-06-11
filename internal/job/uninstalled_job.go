package job

import "e-marketing/internal/service"

type UninstalledJob struct {
	shopService service.ShopService
}

var _ Job = &UninstalledJob{}

func NewUninstalledJob(shopService service.ShopService) *UninstalledJob {
	return &UninstalledJob{
		shopService: shopService,
	}
}

func (job *UninstalledJob) Name() string {
	return "uninstalled"
}

func (job *UninstalledJob) Run() error {
	return nil
}
