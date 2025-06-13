package job

import (
	"context"
	"e-marketing/internal/model"
	"e-marketing/internal/service"
	"time"
)

type NotInstalledJob struct {
	shopService  service.ShopService
	emailService service.EmailService
	timeout      time.Duration
}

var _ Job = &NotInstalledJob{}

func NewUninstalledJob(
	shopService service.ShopService,
	emailService service.EmailService,
) *NotInstalledJob {
	return &NotInstalledJob{
		shopService:  shopService,
		emailService: emailService,
		timeout:      10 * time.Minute,
	}
}

func (job *NotInstalledJob) Name() string {
	return "uninstalled"
}

func (job *NotInstalledJob) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), job.timeout)
	defer cancel()

	// 拿到已发送的偏移量
	offset, err := job.emailService.GetNotInstalledOffet(ctx, "not_installed")
	if err != nil {
		return err
	}

	// 获取 sender 列表
	senders, err := job.emailService.GetSenderListByPurpose(ctx, model.SenderPurposeUniverse)
	if err != nil {
		return err
	}

	scenario, err := job.emailService.GetScenarioByCode(ctx, model.ScenarioCodeNotInstalled)
	if err != nil {
		return err
	}

	triggerRule, err := job.emailService.GetTriggerRuleByScenarioId(ctx, scenario.Id)
	if err != nil {
		return err
	}

	days := triggerRule.Params["days"].([]int)
	for _, day := range days {

	}

	return nil
}
